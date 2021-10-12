package greetings

import ( 
	"errors"	
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" { 
		return "", errors.New("name is required, but missing")
	}
	
	message := fmt.Sprintf("Aloha, %v, welcome to Go! It's a fun language to learn.", name)
	return message, nil
}
