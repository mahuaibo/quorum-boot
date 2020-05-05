package main

import (
	"flag"
	"fmt"
	"github.com/mahuaibo/quorum-boot/boot-cli/http"
	"os"
	"strings"
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

		// genesis content
		genesis, enodes, err := http.GetConsortium(*serverUrl, *consortium)
		if err != nil {
			fmt.Println("err:", err)
			return
		}

		// node dir
		path := *out + "/" + *consortium + "-node"
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Println(err)
		}

		// genesis.json
		err = createFile(path+"/genesis.json", genesis)
		if err != nil {
			fmt.Println(err)
			return
		}

		// static-nodes.json
		err = createFile(path+"/static-nodes.json", enodes)
		if err != nil {
			fmt.Println(err)
			return
		}

		// nodeKey
		err = createFile(path+"/nodeKey", []byte(handlePrivateKey(*privateKey)))
		if err != nil {
			fmt.Println(err)
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

func handlePrivateKey(privateKey string) string {
	privateKey = strings.ToLower(privateKey)
	privateKey = strings.ReplaceAll(privateKey, "0x", "")
	return privateKey
}

func createFile(filename string, content []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	return err
}
