package main

import "fmt"

func main() {
	name := "Ashish Rao"
	tpl := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello world</title>
		</head>
		<body>
		<h1> ` + name + `</h1>
		</body>
		</html>
	`
	fmt.Println(tpl)
}
