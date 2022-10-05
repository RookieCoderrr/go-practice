package main
import (
"bufio"
"fmt"
"net"
"time"
)
//Question1： p is a byte slice of 2048 length.
func main() {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", ":1234")
	if err != nil {
		fmt.Printf("Some error in connection %v ", err)
		return
	}
	for {
		fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
		reader := bufio.NewReader(conn)
		_, err = reader.Read(p)
		if err == nil {
			fmt.Printf("%s\n", p)
		} else {
			fmt.Printf("Some error in reading %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}
	conn.Close()
}
