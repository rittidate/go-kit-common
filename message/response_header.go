package message

type ResponseHeader struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
