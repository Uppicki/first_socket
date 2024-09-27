package wsserviceclient

type wsClient struct {
}

func (client *wsClient) Run() {}

func NewWSClient() IWSClient {
	return &wsClient{}
}
