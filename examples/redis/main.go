package main

import (
	"fmt"

	"github.com/fermyon/spin-go-sdk/redis"
)

func init() {
	// redis.Handle() must be called in the init() function.
	redis.Handle(func(payload []byte) error {
		fmt.Println("Payload::::")
		fmt.Println(string(payload))
		return nil
	})
}

// main function must be included for the compiler but is not executed.
func main() {}
