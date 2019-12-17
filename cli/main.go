/*
	cli interface to encode a given CAPCODE and text message into a POCSAG message
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/kgolding/go-pocsag-encode"
)

func main() {
	binaryMode := flag.Bool("binary", false, "output binary string")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("usage: " + os.Args[0] + ` <capcode> "<message>"`)
		flag.Usage()
		os.Exit(1)
	}

	// Convert first arg to an integer
	id, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	// Second arg is the text message
	message := flag.Arg(1)

	fmt.Printf("CAPCODE = %d, Message = '%s'\n", id, message)

	encodedMsg := pocsagencode.EncodeTransmission(id, message)

	for _, v := range encodedMsg {
		if *binaryMode {
			// Print as a binary bit string
			fmt.Printf("%b", v)
		} else {
			// Print as ASCII hex
			fmt.Printf("%X", v)
		}
	}
	// Print new line
	fmt.Println()
}
