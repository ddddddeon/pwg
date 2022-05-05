package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var chars = "1234567890ABCDEFGHIJKLMNPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

func main() {
	var pwlen = 16
	var err error

	if len(os.Args) >= 2 {
		pwlen, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("First argument must be an int (or nothing): %s\n", err)
			os.Exit(1)
		}
	}

	rand.Seed(time.Now().UnixNano())

	bytes := make([]byte, pwlen)
	for i := 0; i < pwlen; i++ {
		bytes[i] = byte(chars[rand.Int()%len(chars)])
	}

	fmt.Println(string(bytes))
}
