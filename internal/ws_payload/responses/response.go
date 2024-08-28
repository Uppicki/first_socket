package wsresponses

type WSResponse struct {
	ResponseType WSResponseType `json:"type"`
	Response     IWSResponse    `json:"response"`
}
