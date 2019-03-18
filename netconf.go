package networkapi

import (
	"encoding/json"
	"fmt"

	junos "github.com/scottdware/go-junos"
)

type CommitHistory struct {
	User      string `json:"user"`
	Method    string `json:"method"`
	Log       string `json:"log"`
	Comment   string `json:"comment"`
	Timestamp string `json:"timestamp"`
}

// NetworkAPI ... Interface for library connecting over netconf
type NetworkAPI interface {
	Connect() (*junos.Junos, error)
	GetCommitHistory(session *junos.Junos) (string, error)
	GetConfig(session *junos.Junos, format string) (string, error)
	GetInterfaces(session *junos.Junos, format string) (string, error)
	GetLogs(session *junos.Junos) (string, error)
	GetInterfaceEvents(session *junos.Junos) (string, error)
	GetRouterTime(session *junos.Junos) (string, error)
	Close() *junos.Junos
}

//Connect ...
func (c *Client) Connect() (*junos.Junos, error) {

	auth := &junos.AuthMethod{
		Credentials: []string{c.Username, c.Password},
	}

	jnpr, err := junos.NewSession(c.Hostname, auth)
	if err != nil {
		return nil, err
	}
	return jnpr, nil
}

//Close ...
func (c *Client) Close(session *junos.Junos) {
	session.Close()
}

// GetCommitHistory ...
func (c *Client) GetCommitHistory(session *junos.Junos) (string, error) {

	cmtHistory, err := session.CommitHistory()
	if err != nil {
		return "", err
	}

	var cmthtry []CommitHistory

	for _, history := range cmtHistory.Entries {

		data := CommitHistory{
			User:      history.User,
			Method:    history.Method,
			Log:       history.Log,
			Comment:   history.Comment,
			Timestamp: history.Timestamp,
		}

		cmthtry = append(cmthtry, data)
	}

	output, _ := json.Marshal(cmthtry)

	return string(output), nil
}

//GetConfig ...
func (c *Client) GetConfig(session *junos.Junos, format string) (string, error) {

	config, err := session.GetConfig(format)
	if err != nil {
		return "", err
	}

	return config, nil
}

//GetInterfaces ...
func (c *Client) GetInterfaces(session *junos.Junos, format string) (string, error) {

	interfaces, err := session.GetConfig(format, "interfaces")
	if err != nil {
		return "", err
	}

	return interfaces, nil
}

// GetLogs ...
func (c *Client) GetLogs(session *junos.Junos) (string, error) {

	command := fmt.Sprintf("show log messages")

	logs, err := session.Command(command)
	if err != nil {
		return "", err
	}

	return logs, nil
}

// GetInterfaceEvents ...
func (c *Client) GetInterfaceEvents(session *junos.Junos) (string, error) {

	command := fmt.Sprintf("show log intf-events")

	logs, err := session.Command(command)
	if err != nil {
		return "", err
	}

	return logs, nil
}

// GetRouterTime ...
func (c *Client) GetRouterTime(session *junos.Junos) (string, error) {

	rTime, err := session.Command("show system uptime")
	if err != nil {
		return "", err
	}

	return rTime, nil

}
