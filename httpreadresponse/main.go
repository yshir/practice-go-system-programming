package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	res, _ := http.ReadResponse(bufio.NewReader(conn), nil)
	fmt.Println(res.Header)
	fmt.Println("----------------")
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
