package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	special   = "!@#$%^&*()"
)

func main (){
	length := flag.Int("length", 12, "password length (minimum 1)")
	useLower := flag.Bool("lower", true, "include lowercase letters")
	useUpper := flag.Bool("upper", true, "include uppercase letters")
	useNumbers := flag.Bool("numbers", true, "include numbers")
	useSpecial := flag.Bool("special", true, "include special characters")

	flag.Parse()

	if *length < 1 {
		fmt.Println("Error: Password length must be at least 1")
		return
	}

	var charSet strings.Builder
	if *useLower {
		charSet.WriteString(lowercase)
	}
	if *useUpper {
		charSet.WriteString(uppercase)
	}
	if *useNumbers {
		charSet.WriteString(numbers)
	}
	if *useSpecial {
		charSet.WriteString(special)
	}

	if charSet.Len() == 0 {
		fmt.Println("Error: No character types selected")
		return
	}

	charSetStr := charSet.String()

	var password strings.Builder
	for i := 0; i < *length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSetStr))))
		if err != nil {
			fmt.Println("Error generating password:", err)
			return
		}
		password.WriteByte(charSetStr[idx.Int64()])
	}
	fmt.Println(password.String())
}