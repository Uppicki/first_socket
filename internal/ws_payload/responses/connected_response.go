package wsresponses

type ConnectedResponse struct {
	ConnectedType ConnectedResponseType `json:"type"`
	Username      string                `json:"username"`
}
