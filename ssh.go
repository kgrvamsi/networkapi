package networkapi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"strings"

	"golang.org/x/crypto/ssh"
)

type InterfacesInfo struct {
	XMLName              xml.Name              `xml:"rpc-reply"`
	InterfaceInformation InterfaceInformations `xml:"interface-information"`
}

type InterfaceInformations struct {
	XMLName           xml.Name            `xml:"interface-information"`
	PhysicalInterface []PhysicalInterface `xml:"physical-interface"`
	LogicalInterface  []LogicalInterface  `xml:"logical-interface"`
}

type PhysicalInterface struct {
	XMLName     xml.Name `xml:"physical-interface"`
	Name        string   `xml:"name"`
	Adminstatus string   `xml:"admin-status"`
	Operstatus  string   `xml:"oper-status"`
	Description string   `xml:"description"`
}

type LogicalInterface struct {
	XMLName     xml.Name `xml:"logical-interface"`
	Name        string   `xml:"name"`
	Adminstatus string   `xml:"admin-status"`
	Operstatus  string   `xml:"oper-status"`
	Description string   `xml:"description"`
}

type InterfacesList struct {
	Interfacename string `json:"interfacename"`
	Adminstatus   string `json:"adminstatus"`
	Operstatus    string `json:"operstatus"`
	Description   string `json:"description"`
}

type RPCReplyBgp struct {
	XMLName        xml.Name       `xml:"rpc-reply"`
	Bgpinformation BGPInformation `xml:"bgp-information"`
}

type BGPInformation struct {
	XMLName               xml.Name  `xml:"bgp-information"`
	Groupcount            string    `xml:"group-count"`
	Peercount             string    `xml:"peer-count"`
	Downpeercount         string    `xml:"down-peer-count"`
	Unconfiguredpeercount string    `xml:"unconfigured-peer-count"`
	Bgppeer               []Bgppeer `xml:"bgp-peer"`
}

type Bgppeer struct {
	XMLName         xml.Name `xml:"bgp-peer"`
	Peeraddress     string   `xml:"peer-address"`
	Peeras          string   `xml:"peer-as"`
	Inputmessages   string   `xml:"input-messages"`
	Outputmessages  string   `xml:"output-messages"`
	Routequeuecount string   `xml:"route-queue-count"`
	Flapcount       string   `xml:"flap-count"`
	Elapsedtime     string   `xml:"elapsed-time"`
	Peerstate       string   `xml:"peer-state"`
}

type Bgppeers struct {
	Peeraddress string
	Peeras      string
	Peerstate   string
}

// NetworkSSH ... Interface for library connecting over ssh
type NetworkSSH interface {
	ConnectSSH(*ssh.Session, error)
	GetConfigSSH(session *ssh.Session, format string) (string, error)
	GetInterfacesSSH(session *ssh.Session, format string) (string, error)
	GetBGPStatusSSH(session *ssh.Session, format string) (string, error)
	GetLogMessagesSSH(session *ssh.Session, format string) (string, error)
	GetLogMessagesSSH(session *ssh.Session, format string) (string, error)
	GetCommitHistorySSH(session *ssh.Session, format string) (string, error)
	GetLLDPNeighborsSSH(session *ssh.Session, format string) (string, error)
	GetOutputSSH(session *ssh.Session, command string, format string) (string, error)
}

// ConnectSSH ... Establishes session with the device
func (c *Client) ConnectSSH() (*ssh.Session, error) {
	hostname := c.Hostname + ":22"
	config := &ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", hostname, config)
	if err != nil {
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

// GetConfigSSH ... Returns the configuration of device
func (c *Client) GetConfigSSH(session *ssh.Session, format string) (string, error) {
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run("show configuration | display " + format)
	result := stdoutBuf.String()
	return result, nil
}

//GetInterfacesSSH ...Returns the interfaces details of device
func (c *Client) GetInterfacesSSH(session *ssh.Session, format string) (string, error) {

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run("show interfaces descriptions | display " + format)
	result := stdoutBuf.String()

	var Interfaces InterfacesInfo
	xml.Unmarshal([]byte(result), &Interfaces)
	var Interfacesdetails []InterfacesList

	for i := 0; i < len(Interfaces.InterfaceInformation.PhysicalInterface); i++ {
		Interfacedetails := InterfacesList{Interfaces.InterfaceInformation.PhysicalInterface[i].Name, Interfaces.InterfaceInformation.PhysicalInterface[i].Adminstatus,
			Interfaces.InterfaceInformation.PhysicalInterface[i].Operstatus, Interfaces.InterfaceInformation.PhysicalInterface[i].Description}
		Interfacesdetails = append(Interfacesdetails, Interfacedetails)
	}

	for i := 0; i < len(Interfaces.InterfaceInformation.LogicalInterface); i++ {
		Interfacedetails := InterfacesList{Interfaces.InterfaceInformation.LogicalInterface[i].Name, Interfaces.InterfaceInformation.LogicalInterface[i].Adminstatus,
			Interfaces.InterfaceInformation.LogicalInterface[i].Operstatus, Interfaces.InterfaceInformation.LogicalInterface[i].Description}
		Interfacesdetails = append(Interfacesdetails, Interfacedetails)
	}
	output, _ := json.Marshal(Interfacesdetails)
	return string(output), nil
}

// GetBGPStatusSSH ...
func (c *Client) GetBGPStatusSSH(session *ssh.Session, format string) (string, error) {
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run("show bgp summary | display " + format)
	result := stdoutBuf.String()

	var bgppers2 RPCReplyBgp
	xml.Unmarshal([]byte(result), &bgppers2)

	var Bgppeerslist []Bgppeers
	for i := 0; i < len(bgppers2.Bgpinformation.Bgppeer); i++ {
		Bgppeer := Bgppeers{bgppers2.Bgpinformation.Bgppeer[i].Peeraddress, bgppers2.Bgpinformation.Bgppeer[i].Peeras, bgppers2.Bgpinformation.Bgppeer[i].Peerstate}
		Bgppeerslist = append(Bgppeerslist, Bgppeer)
	}
	output, _ := json.Marshal(Bgppeerslist)
	return string(output), nil
}

//GetLogMessagesSSH ...
func (c *Client) GetLogMessagesSSH(session *ssh.Session, format string) (string, error) {
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run("show log messages |no-more| display " + format)
	result := stdoutBuf.String()
	return result, nil
}

//GetSystemUptimeSSH ...
func (c *Client) GetSystemUptimeSSH(session *ssh.Session, format string) (string, error) {
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run("show system uptime | display " + format)
	result := stdoutBuf.String()
	return result, nil
}

//GetCommitHistorySSH ... Returns commit history
func (c *Client) GetCommitHistorySSH(session *ssh.Session, format string) (string, error) {
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run("show system commit |display " + format)
	result := stdoutBuf.String()
	return result, nil
}

//GetLLDPNeighborsSSH ...
func (c *Client) GetLLDPNeighborsSSH(session *ssh.Session, format string) (string, error) {
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run("show lldp neighbors |display " + format)
	result := stdoutBuf.String()
	return result, nil
}

//GetOutputSSH ...Takes command and expected output format as input and returns output in text, JSON or XML based on the output format
func (c *Client) GetOutputSSH(session *ssh.Session, command string, format string) (string, error) {
	if strings.ToLower(format) == "xml" {
		command = command + " | display xml"
	} else if strings.ToLower(format) == "json" {
		command = command + " | display json"
	}
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(command)
	result := stdoutBuf.String()
	return result, nil
}
