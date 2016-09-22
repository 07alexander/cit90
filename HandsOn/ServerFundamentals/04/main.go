package main

import (
	"io"
	"log"
	"net"
	"fmt"
	"bufio"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		serve(c)
		io.WriteString(c, "Hello World")
		c.Close()
	}
}
func serve(c net.Conn){
	scanner := bufio.NewScanner(c)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == ""{
			break
		}
	}
}