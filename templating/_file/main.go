package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Ashish Rao"
	str := fmt.Sprint(
		`<html>
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
		log.Fatal("Error creating a file")
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
