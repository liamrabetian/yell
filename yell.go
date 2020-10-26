package yell

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func FreeSoftware(nationality string) (string, error) {
	if nationality == "" {
		return "", errors.New("No matter who you are, you're human!")
	}
	message := fmt.Sprintf(randomFormat(), nationality)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"I'm %v and Free Software is for everyone!",
		"%vs are entitled to the same rights as the rest of humanity!",
		"%v Lives Matter!",
	}

	return formats[rand.Intn(len(formats))]
}
