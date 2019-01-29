package networkapi

import (
	"io/ioutil"
	"encoding/json"
	"encoding/xml"
	junos "github.com/scottdware/go-junos"
)

type CommitHistory struct {
	User      string `json:"user"`
	Method    string `json:"method"`
	Log       string `json:"log"`
	Comment   string `json:"comment"`
	Timestamp string `json:"timestamp"`
}

type RouterTime struct {
	Multi_routing_engine_results []struct {
		Multi_routing_engine_item []struct {
			Re_name []struct {
				Data string `json:"data"`
			} `json:"re-name"`
			System_uptime_information []struct {
				Attributes struct {
					Xmlns string `json:"xmlns"`
				} `json:"attributes"`
				Current_time []struct {
					Date_time []struct {
						Attributes struct {
							Junos_seconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"date-time"`
				} `json:"current-time"`
				Last_configured_time []struct {
					Date_time []struct {
						Attributes struct {
							Junos_seconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"date-time"`
					Time_length []struct {
						Attributes struct {
							Junos_seconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"time-length"`
					User []struct {
						Data string `json:"data"`
					} `json:"user"`
				} `json:"last-configured-time"`
				Protocols_started_time []struct {
					Date_time []struct {
						Attributes struct {
							Junos_seconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"date-time"`
					Time_length []struct {
						Attributes struct {
							Junos_seconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"time-length"`
				} `json:"protocols-started-time"`
				System_booted_time []struct {
					Date_time []struct {
						Attributes struct {
							Junos_seconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"date-time"`
					Time_length []struct {
						Attributes struct {
							Junos_seconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"time-length"`
				} `json:"system-booted-time"`
				Time_source []struct {
					Data string `json:"data"`
				} `json:"time-source"`
				Uptime_information []struct {
					Active_user_count []struct {
						Attributes struct {
							Junos_format string `json:"junos:format"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"active-user-count"`
					Date_time []struct {
						Attributes struct {
							Junos_seconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"date-time"`
					Load_average_1 []struct {
						Data string `json:"data"`
					} `json:"load-average-1"`
					Load_average_15 []struct {
						Data string `json:"data"`
					} `json:"load-average-15"`
					Load_average_5 []struct {
						Data string `json:"data"`
					} `json:"load-average-5"`
					Up_time []struct {
						Attributes struct {
							Junos_seconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"up-time"`
					User_table []struct{} `json:"user-table"`
				} `json:"uptime-information"`
			} `json:"system-uptime-information"`
		} `json:"multi-routing-engine-item"`
	} `json:"multi-routing-engine-results"`
}

type NetworkAPI interface {
	Connect() (*junos.Junos, error)
	GetCommitHistory(session *junos.Junos) (string, error)
	GetConfig(session *junos.Junos, format string) (string, error)
	GetInterfaces(session *junos.Junos, format string) (string, error)
	GetLogs(session *junos.Junos) (string, error)
	GetRouterTime(session *junos.Junos) (string, error)
}

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

func (c *Client) GetConfig(session *junos.Junos, format string) (string, error) {

	config, err := session.GetConfig(format)
	if err != nil {
		return "", err
	}

	return config, nil
}

func (c *Client) GetInterfaces(session *junos.Junos, format string) (string, error){
	interfaces, err := session.GetConfig(format, "interfaces")
	if err != nil {
		return "", err
	}

	return interfaces, nil
}

func (c *Client) GetLogs(session *junos.Junos) (string, error){

	logs, err := session.Command("show log messages|no-more")
	if err != nil {
		return "", err
	}

	return logs, nil
}

func (c *Client) GetRouterTime(session *junos.Junos) (string, error){

	rTime, err := session.Command("show system uptime|display json")
	if err != nil {
		return "", err
	}

	_data, _ := ioutil.ReadAll(rTime)
	
	var _rTime RouterTime
	
	json.Unmarshal(_data, &_rTime)

	return string(_rTime), nil

}