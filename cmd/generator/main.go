package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/POKTBridge/signer/service"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatalf("can't generate, empty pk provided")
	}
	if os.Args[1] == "" {
		log.Fatalf("can't generate, empty pk provided")
	}

	pass := generatePassword()

	pk, err := service.EncodePK([]byte(os.Args[1]), pass)
	if err != nil {
		log.Fatalf("can't generate, encode pk error")
	}

	fmt.Printf("Encoded pk: %s\n", base64.StdEncoding.EncodeToString(pk))
	fmt.Printf("Signer passcode: %s\n", pass[:32])
	fmt.Printf("Bridge passcode: %s\n", pass[32:])
}

func generatePassword() string {
	charSet := "RW!Vx&6fMHyPYSB1da*4LTm3ki@5c2ptgDzZ9Gq8w7Ke$XNE#s_jvrJuQnFCAUbh"
	return randomStringGenerator(charSet, 64)

}

func randomStringGenerator(charSet string, codeLength int32) string {
	code := ""
	charSetLength := int32(len(charSet))
	for i := int32(0); i < codeLength; i++ {
		index := randomNumber(0, charSetLength)
		code += string(charSet[index])
	}

	return code
}

func randomNumber(min, max int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return min + int32(rand.Intn(int(max-min)))
}
