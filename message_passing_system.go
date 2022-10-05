package main
import (
	"fmt"
	"time"
)
type Vector struct{
	elements []int
}

/*
Question1:
Before changing 100 * time.Millisecond to 200 * time.Millisecond, the number of passes is 8, after changing,
the the number of passes is 6. The event is still driving.

 */
func main() {
	table := make(chan *Vector)
	go player1("Alice", table)
	go player2("Bob", table)
	vector := &Vector{
		elements: []int{0, 0, 0},
	}
	table <- vector // game on; toss the ball
	time.Sleep(1 * time.Second)
	<-table // game over; grab the ball
}
func player1(name string, table chan *Vector) {
	for {
		vector := <-table
		for i, element := range vector.elements{
			vector.elements[i] = element + i +1
		}
		fmt.Println(name)
		for i := 0; i < len(vector.elements); i++ {
			fmt.Println(vector.elements[i])
		}
		time.Sleep(100 * time.Millisecond)
		table <- vector
	}
}
func player2(name string, table chan *Vector) {
	for {
		vector := <-table
		for i, element := range vector.elements{
			vector.elements[i] = element + i +1
		}
		fmt.Println(name)
		for i := 0; i < len(vector.elements); i++ {
			fmt.Println(vector.elements[i])
		}
		time.Sleep(100 * time.Millisecond)
		table <- vector
	}
}