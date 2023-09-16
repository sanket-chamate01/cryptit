package main

import (
	"fmt"
	"github.com/sanket-chamate01/cryptit.git/decrypt"
	"github.com/sanket-chamate01/cryptit.git/encrypt"
)

func main() {
	str := "I Love Eating"
	encr := encrypt.Encryption(str)
	decr := decrypt.Decryption(encr)
	fmt.Printf("Encryption of <%s> : %s\n", str, encr)
	fmt.Printf("Decryption of <%s> : %s", encr, decr)
}
