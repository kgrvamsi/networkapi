package networkapi

import (
	junos "github.com/scottdware/go-junos"
	"encoding/json"
)

type CommitHistory struct {
	User string `json:"user"`
	Method string `json:"method"`
	Log string `json:"log"`
	Comment string `json:"comment"`
	Timestamp string `json:"timestamp"`
}

type NetworkAPI interface {
	Connect() *junos.Junos, error
	GetCommitHistory(session *junos.Junos) string, error
	GetConfig() string, error
}

func(c *Client) Connect() (*junos.Junos, error){

	auth: = &junos.AuthMethod{
		Credentials: []string{c.Username,c.Password}
	}

	jnpr, err := junos.NewSession(c.Hostname, auth)
	if err != nil {
		return nil, err
	}

	defer jnpr.Close()

	return jnpr, nil
}

func (c *Client) GetCommitHistory(session *junos.Junos) (string, error) {
	
	cmtHistory, err := session.CommitHistory()
	if err != nil {
		return nil, err
	}

	var cmthtry []CommitHistory

	for _, history := range cmtHistory.Entries {

		data := CommitHistory{
			User: history.User,
			Method: history.Method,
			Log: history.Log,
			Comment: history.Comment,
			TimeStamp: history.TimeStamp,
		}

		cmthtry = append(cmthtry, data)
	}

	output, _ := json.Marshal(cmthtry)

	return string(output), nil
}

