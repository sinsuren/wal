package main

import (
	"fmt"
	"io"
	"log"
	"wal"
)

func main() {
	walFile, _ := wal.Open(wal.DefaultOptions)

	chunkPosition, _ := walFile.Write([]byte("some data                1"))
	val, _ := walFile.Read(chunkPosition)
	fmt.Println(string(val))

	_, err := walFile.Write([]byte("some data asdfadsfasd 2"))
	if err != nil {
		log.Println(err)
	}
	_, err = walFile.Write([]byte("some data adsfadf3"))
	if err != nil {
		log.Println(err)
	}

	_, err = walFile.Write([]byte("some data asfasfd4"))
	if err != nil {
		log.Println(err)
	}
	_, err = walFile.Write([]byte("some dataadfadsf 5"))
	if err != nil {
		log.Println(err)
	}
	_, err = walFile.Write([]byte("some datafasfa 6"))
	if err != nil {
		log.Println(err)
	}
	_, err = walFile.Write([]byte("some data asdfasdf7"))
	if err != nil {
		log.Println(err)
	}
	_, err = walFile.Write([]byte("some data adfadsfadsfadsfasdfadsfasdfasdfasdfadsfasdfasdfads8"))
	if err != nil {
		log.Println(err)
	}
	_, err = walFile.Write([]byte("some dataadfadfasdfasdfasdfasdfasdfasdfasdf 9"))
	if err != nil {
		log.Println(err)
	}

	reader := walFile.NewReader()
	for {
		val, pos, err := reader.Next()
		if err == io.EOF {
			break
		}
		fmt.Println(string(val))
		fmt.Println(pos)
	}
}
