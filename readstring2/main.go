package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行目
2行目
3行目`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
