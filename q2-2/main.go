package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{"key", "value"})
	writer.Write([]string{"name", "Tanaka"})
	writer.Write([]string{"age", "25"})
	writer.Write([]string{"foo", "bar"})
	writer.Flush()
}
