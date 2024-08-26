package domain

type ChatMessage struct {
	owner   *User
	message string
}
