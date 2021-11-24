package services

import "fmt"

type SmsService struct {
}

//implementation of MessageService interface
func (s *SmsService) Send(message string) bool {
	fmt.Println("Message Sent : ", message)
	return true
}
