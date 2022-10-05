package main

import (
	"fmt"
	"time"
)

type Message struct {
	terminate string
	max int
}
func node(id int, channelinput chan *Message,
	channeloutput chan *Message, Decision chan *Message) {
	State := new(Message)
	//initialize the message
	State.terminate = "No"
	State.max = id
	Counter := 1
	for {
		select {
			//if node receives a message, compare it with its own state
			case message := <- channelinput:
				//if a terminate message is received, the node records it as the final decision, sends it to the next neighbor and return(stop)
				if message.terminate == "Yes" {
					State.terminate = message.terminate
					channeloutput <- State
					Decision <- State
					fmt.Println("Node ", id, " finished.")
					return
				} else {
					switch {
					// if the message is bigger than its state, the node records the message as its own state and sends it to the next neighbor
					case message.max > State.max:
						State = message
						channeloutput <- State
					// if the message is smaller than its state, discard the message.
					case message.max < State.max:
					//if the message is equal to its state, send a terminate message to the next neighbor.
					case message.max == State.max:
						State.terminate = "Yes"
						channeloutput <- State
					}
				}
			//Initialization the message, at the beginning, each node should send a message to his next neighbor
			default :
				if Counter == 1 {
					fmt.Println("Send initial message.")
					channeloutput <- State
				}
		}
		Counter ++
	}
}
func main() {
	// Add 5 channels
	ch12 := make(chan *Message, 1)
	ch23 := make(chan *Message, 1)
	ch34 := make(chan *Message, 1)
	ch45 := make(chan *Message, 1)
	ch51 := make(chan *Message, 1)
	// add 5 decision channels for each node
	decisionCh1 :=  make(chan *Message, 1)
	decisionCh2 :=  make(chan *Message, 1)
	decisionCh3 :=  make(chan *Message, 1)
	decisionCh4 :=  make(chan *Message, 1)
	decisionCh5 :=  make(chan *Message, 1)
	// set 5 nodes in the graph
	go node(9, ch51, ch12, decisionCh1)
	go node(50, ch12, ch23, decisionCh2)
	go node(82, ch23, ch34, decisionCh3)
	go node(69, ch34, ch45, decisionCh4)
	go node(45, ch45, ch51, decisionCh5)
	// sleep for a while to wait for all of the nodes finish
	time.Sleep(1000* time.Millisecond)
	// choose the fastest decision as the result(every decision channel has the same result, and we choose the fastest one here)
	select {
		case Input1 := <- decisionCh1:
			fmt.Println("The leader is Node",Input1.max)
		case Input1 := <- decisionCh2:
			fmt.Println("The leader is Node",Input1.max)
		case Input1 := <- decisionCh3:
			fmt.Println("The leader is Node",Input1.max)
		case Input1 := <- decisionCh4:
			fmt.Println("The leader is Node",Input1.max)
		case Input1 := <- decisionCh5:
			fmt.Println("The leader is Node",Input1.max)
	}

}
