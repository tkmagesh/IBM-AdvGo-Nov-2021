package services

type MessageService interface {
	Send(message string) bool
}
