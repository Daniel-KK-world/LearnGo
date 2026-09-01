package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

//How do we run the code in our projects? We do go run {filename}

//What does "package mian" mean in Go?
//Package is like a project as a whole, and main is the entry point of the project. It is where the program starts executing.
//Every first line of a Go file must declare the package it belongs to. The main package is a special package that tells the Go compiler that this is an executable program, rather than a library.

//What in the world is import "fmt"?
//The import statement is used to include external packages in our Go program. "fmt" is a standard library package that provides functions for formatting and printing text. In other words, it is giving our package acess to the fmt package, which allows us to use its functions, such as Println, to print output to the console.

//What is func?
//func is the keyword used to define a function in Go.

//How is the file organized?
//Go files are organized with a package declaration at the top, followed by import statements, and then the function definitions.
