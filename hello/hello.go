package main

import (
    "fmt"
    "log"

    "rsc.io/quote"

    "example.com/greetings" // "../greetings" ? currently ran this, as suggested by tutorial
							// must run `$ go mod edit -replace example.com/greetings=../greetings` since example.com/greetings is not published. Alt, would use dependency/bundler equivalent
							// go.mod file updated via above command :)
)

// func main() {
//     fmt.Println("Hello, World!")
// }

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // fmt.Println(quote.Go())
	greeting, err := greetings.Hello("Amanda")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

	pithyQuote := quote.Go()
	fmt.Println(greeting + " " + pithyQuote)
}
