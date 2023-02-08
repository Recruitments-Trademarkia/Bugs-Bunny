package schemas

type WsResponse struct {
	Output      string `json:"output"`
	Error       string `json:"error"`
	ServerError string `json:"server_error"`
}
