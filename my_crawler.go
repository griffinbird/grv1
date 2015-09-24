package main

import (
	"flag" // 'flag' helps you parse command line arguments
	"fmt"
	"os" // 'os' gives you access to system calls
)

func main() {
	flag.Parse()        // This stuff with 'flag' converts the command line
	args := flag.Args() // arguments to a new variable named 'args'
	// The ':=' form means "this is a brand new variable" and
	// a '=' here would throw an error.
	// At this point 'args' will be either an empty list
	// (if no command line arguments were provided) or it'll
	// contain some string values.
	if len(args) < 1 {
		fmt.Println("Please specify start page") // if a starting page wasn't provided as an argument
		os.Exit(1)                               // show a message and exit.
	} // Note that 'main' doesn't return anything.
}
