package message

type RequestHeader struct {
	AppVersion      string   `json:"appVersion" bson:"appVersion"`
	AppID           string   `json:"appId" bson:"appId"`
	SessionID       string   `json:"sessionId" bson:"sessionId"`
	RequestUniqueID string   `json:"requestUniqueId" bson:"requestUniqueId"`
	RequestDateTime string   `json:"requestDateTime" bson:"requestDateTime"`
	CorrID          string   `json:"corrId" bson:"corrId"`
	Language        Language `json:"language" bson:"language"`
}
