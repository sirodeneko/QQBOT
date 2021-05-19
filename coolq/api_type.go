package coolq

type CallType string

const (
	Http      CallType = "http"
	Websocket CallType = "websocket"
)

const (
	SendMSG              Api = "send_msg"                //发送消息
	DeleteMsg            Api = "delete_msg"              //撤回消息
	SetGroupKick         Api = "set_group_kick"          //踢人
	SetGroupBan          Api = "set_group_ban"           //禁言
	SetGroupAnonymousBan Api = "set_group_anonymous_ban" //禁匿名用户的言
	SetGroupWholeBan     Api = "set_group_whole_ban"     //全群禁言
	SetGroupCard         Api = "set_group_card"          //设置群名片（群备注）
	SetGroupSpecialTitle Api = "set_group_special_title" //设置群头衔
	SendGroupNotice      Api = "_send_group_notice"      //发送群公告

)

type SendMSGParams struct {
	MessageType MessageType `json:"message_type"` // 消息类型
	UserId      int64       `json:"user_id"`      // 对方id
	GroupId     int64       `json:"group_id"`     // 群组id
	Message     string      `json:"message"`      // 消息
	AutoEscape  bool        `json:"auto_escape"`  // 是否当成纯文本
}

type DeleteMsgParams struct {
	MessageID int64 `json:"message_id"`
}

type SetGroupKickParams struct {
	GroupId          int64 `json:"group_id"`
	UserId           int64 `json:"user_id"`
	RejectAddRequest bool  `json:"reject_add_request"` //是否拒绝此人的加群请求
}

type SetGroupBanParams struct {
	GroupId  int64 `json:"group_id"`
	UserId   int64 `json:"user_id"`
	Duration int64 `json:"duration"` // 禁言时长 单位秒  0表示不禁言
}

type SetGroupAnonymousBanParams struct {
	GroupId       int64  `json:"group_id"`
	AnonymousFlag string `json:"anonymous_flag"` //匿名用户的flag
	Duration      int64  `json:"duration"`       // 禁言时长 单位秒  0表示不禁言
}

type SetGroupWholeBanParams struct {
	GroupId int64 `json:"group_id"`
	Enable  bool  `json:"enable"` //是否禁言
}

type SetGroupCardParams struct {
	GroupId int64  `json:"group_id"`
	UserId  int64  `json:"user_id"`
	Card    string `json:"card"` //为空则表示删除
}

type SetGroupSpecialTitleParams struct {
	GroupId      int64  `json:"group_id"`
	UserId       int64  `json:"user_id"`
	SpecialTitle string `json:"special_title"` //专属头衔, 不填或空字符串表示删除专属头衔
	Duration     int64  `json:"duration"`      // 专属头衔有效期, 单位秒, -1 表示永久, 不过此项似乎没有效果, 可能是只有某些特殊的时间长度有效, 有待测试
}

type SendGroupNoticeParams struct {
	GroupId int64  `json:"group_id"`
	Content string `json:"content"`
}
