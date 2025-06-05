package hello

import (
	"errors"
	"fmt"
	"log"
)

func Hello(name string) (string, error) {
	log.Println(name)
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hello %s", name)
	return message, nil
}
