package types

type PostNotificationTaskPayload struct {
	Question           string
	ConnectionType     string
	PostID             string
	UserID             string
	ConnectionUserID   string
	IsInternalQuestion bool
	TopicRounds        [][]string
}

type ChatNotificationTaskPayload struct {
	PostID     string
	FromUserID string
	ToUserID   string
	Message    string
	MessageID  string
}

type FriendNotificationTaskPayload struct {
	ConnectionID   string
	SenderUserID   string
	ReceiverUserID string
}

// A list of task types.
const (
	TypePostNotification   = "post:deliver"
	TypeChatNotification   = "chat:deliver"
	TypeFriendNotification = "friend:deliver"
	TypeApprovedFriend     = "approvedFriendRequest:deliver"
	TypeSystemNotification = "system:deliver"
)
