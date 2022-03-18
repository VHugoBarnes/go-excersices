package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {

	c := make(chan string, 2)

	// Add str's to channel
	c <- "Message 1"
	c <- "Message 2"

	// With len we check how many routines are in the channel
	// With cap we check the maximmum capacity of the channel
	fmt.Println(len(c), cap(c))

	// Loop through channel with:
	// range & close
	close(c)
	for message := range c {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)

	go message("Hola", email1)
	go message("Mundo", email2)

	// Select
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}
