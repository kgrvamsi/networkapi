package networkapi

//Client Initialize the Constructor Variables
type Client struct {
	Hostname string
	Username string
	Password string
}

//NetworkClient Initialize the Constructor
func NetworkClient(hostname string, username string, password string) *Client {

	return &Client{
		Hostname: hostname,
		Username: username,
		Password: password,
	}
}
