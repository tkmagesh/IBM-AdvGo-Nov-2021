package services

type MessageProcessor struct {
	messageService MessageService
}

func NewMessageProcessor(messageService MessageService) *MessageProcessor {
	return &MessageProcessor{messageService}
}

func (mp *MessageProcessor) Process(message string) bool {
	return mp.messageService.Send(message)
}
