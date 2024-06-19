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

// A list of task types.
const (
	TypePostNotification   = "post:deliver"
	TypeChatNotification   = "chat:deliver"
	TypeSystemNotification = "system:deliver"
)
