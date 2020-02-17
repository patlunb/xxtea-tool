package main

import (
	"encoding/hex"
	"fmt"
	"flag"
	"strings"
	"github.com/hillu/go-xxtea"
)


func main(){

	var keyString, inputString string
	var decodeVar, versionVar, silentVar bool

	VERSION := "1.0"

	//Making some run-flags
	flag.StringVar(&keyString,"key", "0123456789abcdef", "Key for encoding / decoding")
	flag.StringVar(&inputString,"input", "", "String for encoding or hex string for decoding")

	flag.BoolVar(&decodeVar, "d", false,"Decode string bool flag. If not set, encode operation will be performed")
	flag.BoolVar(&versionVar, "v", false, "Show version")
	flag.BoolVar(&silentVar, "s", false, "Suspend all info messages, outputs only a result of conversion. Good for batch operation or external usage.")

	flag.Parse()

	inputString = strings.TrimSpace(inputString)
	keyString = strings.TrimSpace(keyString)
	
	// A bunch of error handling and input parsing

	if (versionVar) {
		fmt.Println(VERSION)
		return
	}
	
	if (len(keyString) > 16) {
		fmt.Println("Error: Key can't be larger 16 characters")
		return
	}

	if (len(keyString) == 0 ) {
		fmt.Println("Error: Can't operate with empty key")
		return
	}

	if (len(inputString) == 0) {
		fmt.Println("Error: Need to specify [-input] flag")
		return
	}


	//main operation starts here

	if (decodeVar){
		decrypt(keyString, inputString, silentVar)
	} else {
		encrypt(keyString, inputString, silentVar)
	}

}




func decrypt(key, encoded string, suspendMessages bool) (string) {

	//construct key, add some zeroes if not 16 bytes
	k := []byte(key)
	for (len(k) % 16) != 0 {
	  k = append(k, 0)
	}

	//print some debug info
	if (!suspendMessages){
		fmt.Print("Decrypting: ")
		fmt.Println(encoded)
		fmt.Print("Using key: ")
		fmt.Println(key);
	}

	//create cipher based on key
	cipher, err := xxtea.NewCipher(k)

	//panic if anything goes wrong
	if err != nil {
		panic(err)
	}

	// hex string -> byte array
	crypted, err:= hex.DecodeString(encoded)

	//creating empty output array
	plain := make([]byte, len(crypted))

	//decryption
	cipher.Decrypt(plain, crypted)


	//converting to string
	outp := string(plain)

	//trim spaces
	outp = strings.TrimSpace(outp)

	//remove trailing zeroes
	for (outp [len(outp) -1]) == 0x00 {
		outp = strings.TrimSuffix(outp, "\x00")
	  }

	//never hurts to trim spaces once more
	outp = strings.TrimSpace(outp)
	
	//if not in silent mode - print some info
	if (!suspendMessages){
		fmt.Println("Decryption Done. Result:")
	}

	//print result
	fmt.Println(outp)

	//bail out, return result if needed
	return outp

}


func encrypt(key, toencode string, suspendMessages bool) (string) {

	//construct key, add some zeroes if not 16 bytes
	k := []byte(key)
	for (len(k) % 16) != 0 {
	  k = append(k, 0)
	}

	//create cipher based on key
	cipher, err := xxtea.NewCipher(k)

	//panic if anything goes wrong
	if err != nil {
		panic(err)
	}

	//if not in silent mode, print some info
	if (!suspendMessages){
		fmt.Print("Encrypting: ")
		fmt.Println(toencode)
		fmt.Print("Using key: ")
		fmt.Println(key);
	}


	//convert input string to byte array
	plain := []byte(toencode)

	//add some zeroes to match cipher block size
	for (len(plain) % cipher.BlockSize()) != 0 {
		plain = append(plain, 0)
	  }

	//create empty array for output
	crypted := make([]byte, len(plain))

	//main encryption job
	cipher.Encrypt(crypted, plain)

	//let's make some nice hex representation
	dst := make([]byte, hex.EncodedLen(len(crypted)))
	hex.Encode(dst, crypted)
	outp := string(dst)
	outp = strings.ToUpper(outp)
	outp = strings.TrimSpace(outp)

	  //print hoorray message
	if (!suspendMessages){
		fmt.Println("Encryption Done. Result:")
	}
 
	//print result and return
	fmt.Println(outp)
	return outp
}
