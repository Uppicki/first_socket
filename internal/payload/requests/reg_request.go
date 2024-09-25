package requests

type RegRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
