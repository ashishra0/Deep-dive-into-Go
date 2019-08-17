package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]      // The argument you pass when running the program
	fmt.Println(os.Args[0]) // The entire cwd of the program
	fmt.Println(os.Args[1])
	str := fmt.Sprint(`
	<html>
		<head>
			<title>Hello world</title>
		</head>
		<body>
			<h1> ` + name + ` </h1>
		</body>
	</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating html file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}
