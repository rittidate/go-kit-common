package message

type Response struct {
	Header ResponseHeader `json:"header"`
	Data   interface{}    `json:"data"`
}
