package main

//import encrypt package within a module
import (
	"fmt"
	"github.com/MIna-Maher/myencryptor/decrypt"
	"github.com/MIna-Maher/myencryptor/encrypt"
)

func main() {
	string := encrypt.StringEncryptor("Hi Mina Maher!")
	fmt.Println(string)
	fmt.Println(decrypt.StringDecryptor(string))
}
