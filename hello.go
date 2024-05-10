package main

import (
	"fmt"
	"log"
    "golang.org/x/example/hello/reverse"
	"example.com/greetings"
)
func main() {

    //String reverse
    fmt.Println(reverse.String("Hello"))
    //Set the properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and the line number.
    log.SetPrefix("Greetings: ")
    log.SetFlags(0)

    //a slice of names
    names := []string{"Gladys","Ekavir","Samantha","Terrian"}

    //Request greetings for the names 
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map 
    // of messages to the console.
    fmt.Println(messages)
     
}
