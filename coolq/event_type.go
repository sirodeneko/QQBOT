package coolq

type Event string

const (
	PrivateMessageEvent Event = "privateMessageEvent" // 私聊消息
	GroupMessageEvent   Event = "groupMessageEvent"   // 群消息
	GroupIncreaseEvent  Event = "groupIncrease"       //群成员增加
)

type MessageType string

const (
	Private MessageType = "private"
	Group   MessageType = "group"
)

type NoticeType string

const (
	GroupIncreaseNotice NoticeType = "group_increase"
)

type Sender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`

	// 以下属性为群消息才有意义，并且并不保证一定存在
	Card string `json:"card"` //群名片
	Role string `json:"role"` //角色 owner,admin,member
}

type PrivateMessage struct {
	PostType    string      `json:"post_type"`
	MessageType MessageType `json:"message_type"`
	SubType     string      `json:"sub_type"`
	MessageId   int         `json:"message_id"`
	UserId      int64       `json:"user_id"`
	TargetId    int64       `json:"target_id"`
	Message     string      `json:"message"`
	RawMessage  string      `json:"raw_message"`
	SelfId      int64       `json:"self_id"`
	Time        int64       `json:"time"`
	Sender      Sender      `json:"sender"`
}

type Anonymous struct {
	Flag string `json:"flag"`
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GroupMessage struct {
	PostType    string      `json:"post_type"`
	MessageType MessageType `json:"message_type"`
	SubType     string      `json:"sub_type"` //normal正常消息 anonymous 匿名消息 notice 系统提示
	MessageId   int         `json:"message_id"`
	MessageSeq  int64       `json:"message_seq"` //消息在cqhttp中的id
	GroupId     int64       `json:"group_id"`
	UserId      int64       `json:"user_id"`
	Message     string      `json:"message"`
	RawMessage  string      `json:"raw_message"`
	Font        int         `json:"font"` //字体？
	SelfId      int64       `json:"self_id"`
	Time        int64       `json:"time"`
	Sender      Sender      `json:"sender"`
	Anonymous   *Anonymous  `json:"anonymous"`
}

type GroupIncrease struct {
	PostType   string     `json:"post_type"`
	NoticeType NoticeType `json:"notice_type"`
	GroupId    int64      `json:"group_id"`
	OperatorId int64      `json:"operator_id"` //操作者 QQ 号
	SelfId     int64      `json:"self_id"`
	SubType    string     `json:"sub_type"` //approve、invite 分别表示管理员已同意入群、管理员邀请入群
	Time       int64      `json:"time"`
	UserId     int64      `json:"user_id"`
}
