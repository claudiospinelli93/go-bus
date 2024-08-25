package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/claudiospinelli93/go-bus/helper"
)

func main() {

	servicebusConnectionString, shouldReturn := setServiceBusConnectionString()
	if shouldReturn {
		return
	}

	queueOrTopicSource := helper.PromptUser("Enter source: ")
	queueOrTopicDest := helper.PromptUser("Enter destination: ")

	maxMessagesCount, err := helper.PromptUserInt("Max receiver mensage (default 100): ", 100)
	if err != nil {
		return
	}

	senderParallelCount, err := helper.PromptUserInt("Parallel sender number (default 50): ", 50)
	if err != nil {
		return
	}

	client := getClient(servicebusConnectionString)
	defer client.Close(context.TODO())

	receiverSource := createReceiver(queueOrTopicSource, client)
	defer receiverSource.Close(context.TODO())

	sender := createSender(queueOrTopicDest, client)
	defer sender.Close(context.TODO())

	sem := make(chan struct{}, senderParallelCount)
	interation := 0

	for {
		interation++

		messagesPeek := existsMessages(receiverSource)
		if len(messagesPeek) == 0 {
			fmt.Println("No messages found!")
			break
		}

		messages := getReceiveMessages(maxMessagesCount, receiverSource)

		if len(messages) == 0 {
			fmt.Println("No messages found!")
			break
		}

		var wg sync.WaitGroup
		for _, message := range messages {
			wg.Add(1)
			sem <- struct{}{}
			go func(msg *azservicebus.ReceivedMessage) {
				defer wg.Done()
				defer func() { <-sem }()
				body := message.Body
				sendMessage(string(body), sender)
				err := receiverSource.CompleteMessage(context.TODO(), message, nil)
				if err != nil {
					panic(err)
				}
			}(message)
		}
		wg.Wait()
	}

	fmt.Println("Message sent successfully!")
}

func setServiceBusConnectionString() (string, bool) {
	var enterServicebusConnectionString string
	var servicebusConnectionString string

	servicebusConnectionString = helper.GetEnv("SERVICE_BUS_CONNECTION_STRING", "empty")

	enterServicebusConnectionString = helper.PromptUser("Enter connection string (default " + servicebusConnectionString + "): ")
	if enterServicebusConnectionString != "" {
		servicebusConnectionString = enterServicebusConnectionString
	}

	if servicebusConnectionString == "empty" {
		fmt.Println("SERVICE_BUS_CONNECTION_STRING is required.")
		return "", true
	}
	return servicebusConnectionString, false
}

func getClient(servicebusConnectionString string) *azservicebus.Client {
	client, err := azservicebus.NewClientFromConnectionString(servicebusConnectionString, nil)
	if err != nil {
		panic(err)
	}
	return client
}

func sendMessage(message string, sender *azservicebus.Sender) {
	sbMessage := &azservicebus.Message{
		Body: []byte(message),
	}
	err := sender.SendMessage(context.TODO(), sbMessage, nil)
	if err != nil {
		panic(err)
	}
}

func createReceiver(queueOrTopic string, client *azservicebus.Client) *azservicebus.Receiver {
	receiver, err := client.NewReceiverForQueue(queueOrTopic, nil)
	if err != nil {
		panic(err)
	}
	return receiver
}

func createSender(queueOrTopic string, client *azservicebus.Client) *azservicebus.Sender {
	sender, err := client.NewSender(queueOrTopic, nil)
	if err != nil {
		panic(err)
	}
	return sender
}

func getReceiveMessages(count int, receiver *azservicebus.Receiver) []*azservicebus.ReceivedMessage {
	messages, err := receiver.ReceiveMessages(context.TODO(), count, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	return messages
}

func existsMessages(receiver *azservicebus.Receiver) []*azservicebus.ReceivedMessage {
	messages, err := receiver.PeekMessages(context.TODO(), 1, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	return messages
}
