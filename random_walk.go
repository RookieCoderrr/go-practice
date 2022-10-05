package main
import (
	"fmt"
	"math/rand"
	"time"
)

//Token represent the message passed in the graph
type Token struct {
	data string
}
/*
	Question2: 1/N
 */

// return an integer from a given range
func getRandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
// the node with 3 neighbors
func nodeWith3Neighbor(name string, channel1 chan *Token, channel2 chan *Token, channel3 chan *Token) {
	for {
		// execute when a token received from a certain channel
		select {
		case token := <-channel1:
			fmt.Println(name," received", token.data)
			random := getRandomInt(3) + 1
			// randomly sent the token to its neighbor
			switch random {
			case 1:
				channel1 <- token
			case 2:
				channel2 <- token
			case 3:
				channel3 <- token
			}
		case token := <-channel2:
			fmt.Println(name," received", token.data)
			// randomly sent the token to its neighbor
			random := getRandomInt(3) + 1
			switch random {
			case 1:
				channel1 <- token
			case 2:
				channel2 <- token
			case 3:
				channel3 <- token
			}
		case token := <-channel3:
			fmt.Println(name," received", token.data)
			// randomly sent the token to its neighbor
			random := getRandomInt(3) + 1
			switch random {
			case 1:
				channel1 <- token
			case 2:
				channel2 <- token
			case 3:
				channel3 <- token
			}
		}
	}
}
// the node with 2 neighbors
func nodeWith2Neighbor(name string, channel1 chan *Token, channel2 chan *Token) {
	for {
		// execute when a token received from a certain channel
		select {
		case token := <-channel1:
			fmt.Println(name," received", token.data)
			// randomly sent the token to its neighbor
			random := getRandomInt(2) + 1
			switch random {
			case 1:
				channel1 <- token
			case 2:
				channel2 <- token
			}
		case token := <-channel2:
			fmt.Println(name," received", token.data)
			// randomly sent the token to its neighbor
			random := getRandomInt(2) + 1
			switch random {
			case 1:
				channel1 <- token
			case 2:
				channel2 <- token
			}
		}
	}
}
// the node with only 1 neighbor
func nodeWith1Neighbor(name string, channel1 chan *Token) {
	for {
		// Only one channel connected to this node to send or receive token
		token := <-channel1
		fmt.Println(name, " received", token.data)
		channel1 <- token
	}
}

func main() {
	// 6 channel represents 6 edges in the graph
	ch1 := make(chan *Token)
	ch2 := make(chan *Token)
	ch3 := make(chan *Token)
	ch4 := make(chan *Token)
	ch5 := make(chan *Token)
	ch6 := make(chan *Token)
	// 5 thread represent 5 nodes linked to the 6 edges in the graph
	go nodeWith2Neighbor("Node2_a", ch1, ch2)
	go nodeWith3Neighbor("Node3_b", ch1, ch3, ch4)
	go nodeWith3Neighbor("Node3_c", ch2, ch3, ch5)
	go nodeWith3Neighbor("Node3_d", ch4, ch5, ch6)
	go nodeWith1Neighbor("Node1_e", ch6)
	//initialize the token carried "secret" as its message
	token := &Token{
		data: "secret",
	}
	fmt.Println("Start")
	// start the token in the channel 2
	ch2 <- token
	// set a process time
	time.Sleep(5 * time.Second)
	fmt.Println("Ops! Stop")
}