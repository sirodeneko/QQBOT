package coolq

type Sender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
}

type PrivateMessage struct {
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int    `json:"message_id"`
	UserId      int64  `json:"user_id"`
	TargetId    int64  `json:"target_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	SelfId      int64  `json:"self_id"`
	Time        int64  `json:"time"`
	Sender      Sender `json:"sender"`
}
