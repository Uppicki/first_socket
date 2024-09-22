package domain

type Chat struct {
	ID       string
	users    []*User
	messages []*ChatMessage
}

func NewChat(user1, user2 *User) *Chat {
	return &Chat{
		ID:       "",
		users:    make([]*User, 2),
		messages: make([]*ChatMessage, 0),
	}
}
