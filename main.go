package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	user2 "os/user"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello, and welcome to my QotaqScript interpreted turing complete language" + user.Username)
	repl.Start(os.Stdin, os.Stdout)
}

// go feature if inside of the function with variable for example x := 42 we will return &42, variable will be allocated on heap memory
