package services

import "testing"

type MockMessageService struct {
	sendInvoked bool
	sentMessage string
	returnValue bool
}

//implementation of the MessageService interface
func (m *MockMessageService) Send(message string) bool {
	m.sendInvoked = true
	m.sentMessage = message
	m.returnValue = true
	return m.returnValue
}

func Test_MessageProcessor_Sends_Message(t *testing.T) {
	//Arrange
	mockMessageService := new(MockMessageService)
	messageProcessor := NewMessageProcessor(mockMessageService)

	//Act
	messageProcessor.Process("Hello World")

	//Assert
	if !mockMessageService.sendInvoked {
		t.Error("Expected Send to be invoked on the MessageService")
	}
}

func Test_MessageProcessor_Sends_The_Given_Message(t *testing.T) {
	//Arrange
	mockMessageService := new(MockMessageService)
	messageProcessor := NewMessageProcessor(mockMessageService)
	msg := "Hello World"

	//Act
	messageProcessor.Process(msg)

	//Assert
	if mockMessageService.sentMessage != msg {
		t.Errorf("Expected message to be %s but was %s", msg, mockMessageService.sentMessage)
	}
}

func Test_MessageProcessor_Returns_True_When_Message_Is_Sent(t *testing.T) {
	//Arrange
	mockMessageService := new(MockMessageService)
	messageProcessor := NewMessageProcessor(mockMessageService)
	msg := "Hello World"

	//Act
	result := messageProcessor.Process(msg)

	//Assert
	if result != mockMessageService.returnValue {
		t.Errorf("Expected result to be %t but was %t", mockMessageService.returnValue, result)
	}
}
