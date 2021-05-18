package coolq

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
	SubType     string      `json:"sub_type"`
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
