package main

import (
	"github.com/Ubbo-Sathla/mfa/pkg/mfa"
)

func main() {
	err := mfa.LoadConfig()
	if err != nil {
		return
	}
	c := mfa.GetConfig()
	for _, j := range c {
		j.Display()
	}
}
