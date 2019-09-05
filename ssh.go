package networkapi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

// NetworkSSH ... Interface for library connecting over ssh
type NetworkSSH interface {
	ConnectSSH(*ssh.Session, error)
	GetConfigSSH(session *ssh.Session, format string) (string, error)
	GetInterfacesSSH(session *ssh.Session, format string) (string, error)
	GetBGPStatusSSH(session *ssh.Session, format string) (string, error)
	GetLogMessagesSSH(session *ssh.Session) (string, error)
	GetCommitHistorySSH(session *ssh.Session, port string) (string, error)
	GetLLDPNeighborsSSH(session *ssh.Session, format string) (string, error)
	GetOutputSSH(session *ssh.Session, command string, format string) (string, error)
	CloseSSH(session *ssh.Session)
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

// CloseSSH ...
func (c *Client) CloseSSH(session *ssh.Session) {
	session.Close()
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
func (c *Client) GetLogMessagesSSH(session *ssh.Session) (string, error) {
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	command := fmt.Sprintf("show log messages")
	session.Run(command)
	result := stdoutBuf.String()
	return result, nil
}

//GetSystemUptimeSSH ...
func (c *Client) GetSystemUptimeSSH(session *ssh.Session, format string) (string, error) {
	var (
		stdoutBuf   bytes.Buffer
		result      string
		_routerTime RouterTime
	)
	session.Stdout = &stdoutBuf
	if format == "json" {
		session.Run("show system uptime | display " + format)
		result := stdoutBuf.Bytes()
		json.Unmarshal(result, &_routerTime)

		_result := RouterTimeRes{Currenttime: _routerTime.MultiRoutingEngineResults[0].MultiRoutingEngineItem[0].SystemUptimeInformation[0].Currenttime[0].Datetime[0].Data,
			LastConfiguredTime: _routerTime.MultiRoutingEngineResults[0].MultiRoutingEngineItem[0].SystemUptimeInformation[0].LastConfiguredTime[0].DateTime[0].Data,
			SystemBootedTime:   _routerTime.MultiRoutingEngineResults[0].MultiRoutingEngineItem[0].SystemUptimeInformation[0].SystemBootedTime[0].DateTime[0].Data}

		output, _ := json.Marshal(_result)
		return string(output), nil
	} else {
		session.Run("show system uptime | display " + format)
		result := stdoutBuf.String()
		return result, nil
	}
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
