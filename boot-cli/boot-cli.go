package main

import (
	"flag"
	"fmt"
	"github.com/mahuaibo/quorum-boot/boot-cli/action"
)

var (
	serverUrl  *string
	consortium *string
	privateKey *string
	out        *string
)

func init() {
	serverUrl = flag.String("server_url", "http://127.0.0.1:8080", "http server url.")
	consortium = flag.String("consortium", "", "consortium id.")
	privateKey = flag.String("privateKey", "", "privateKey.")
	out = flag.String("out", "./", "out path.")
	flag.Parse()
}

func main() {
	cmd := flag.Arg(0)
	if cmd == "init" {
		if *consortium == "" {
			fmt.Println("consortium id err")
			return
		}
		if *privateKey == "" {
			fmt.Println("privateKey err")
			return
		}

		err := action.Boot(*serverUrl, *consortium, *privateKey, *out)
		if err != nil {
			fmt.Println("err:", err)
			return
		}

	} else if cmd == "help" {
		fmt.Println(`boot-cli init	
	-publicKey
	-consortium`)
	} else {
		fmt.Println("help")
	}
}
