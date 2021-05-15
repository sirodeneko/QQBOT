package coolq

type CallType string

const (
	Http      CallType = "http"
	Websocket CallType = "websocket"
)

type MessageType string

const (
	Private MessageType = "private"
	Group   MessageType = "group"
)
const (
	SendMSG Api = "send_msg"
)

type sendMSGParams struct {
	MessageType MessageType `json:"message_type"` // 消息类型
	UserId      int64       `json:"user_id"`      // 对方id
	GroupId     string      `json:"group_id"`     // 群组id
	Message     string      `json:"message"`      // 消息
	AutoEscape  string      `json:"boolean"`      // 是否当成纯文本
}