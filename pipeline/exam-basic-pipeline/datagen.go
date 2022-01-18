package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func generateData() <-chan CusUUID {
	c := make(chan CusUUID)
	const filePath = "guids.txt"
	go func() {
		file, _ := os.Open(filePath)
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			line = strings.TrimSuffix(line, "\n")
			uuid := CusUUID{Id: line}
			log.Println(line)

			if err != nil {
				continue
			}

			c <- uuid
		}

		close(c)
	}()
	return c
}
