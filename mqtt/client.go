package mqtt

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	Broker     string
	Id         string
	TopicChat  string
	TopicImage string
}

const (
	broker         = "tcp://192.168.2.12:1883"
	clientID       = "smart-glass"
	topicMsg       = "smart-glass/testo"
	topicDetection = "smart-glass/detection"
)

var mqttMsgChan = make(chan mqtt.Message)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	mqttMsgChan <- msg
}

func processMsg(ctx context.Context, input <-chan mqtt.Message, chmsg chan string, chdetection chan string) chan mqtt.Message {
	out := make(chan mqtt.Message)
	go func() {
		defer close(out)
		for {
			select {
			case msg, ok := <-input:
				if !ok {
					return
				}

				fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
				if msg.Topic() == topicMsg {
					chmsg <- string(msg.Payload())
				} else if msg.Topic() == topicDetection {
					chdetection <- string(msg.Payload())
				}
				out <- msg
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to MQTT Broker")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func NewReciever(chmsg chan string, chdetection chan string) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername("mqtt")
	opts.SetPassword("Paolomanfe95")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		finalChan := processMsg(ctx, mqttMsgChan, chmsg, chdetection)
		for range finalChan {
			// just consuming these for now
		}
	}()

	// Subscribe to the topic
	token := client.Subscribe(topicMsg, 1, nil)

	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topicMsg)
	token = client.Subscribe(topicDetection, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topicMsg)

	// Wait for interrupt signal to gracefully shutdown the subscriber
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	// Cancel the context to signal the goroutine to stop
	cancel()

	// Unsubscribe and disconnect
	fmt.Println("Unsubscribing and disconnecting...")
	client.Unsubscribe(topicMsg)
	client.Unsubscribe(topicDetection)
	client.Disconnect(250)

	// Wait for the goroutine to finish
	wg.Wait()
	fmt.Println("Goroutine terminated, exiting...")
}
