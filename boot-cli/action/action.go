package action

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mahuaibo/quorum-boot/boot-cli/common"
	"os"
)

func Boot(serverUrl, consortium, privateKey, out string) error {
	value, err := common.HttpGet(serverUrl, "/consortiums/"+consortium+"/boot")
	if err != nil {
		return err
	}
	response := struct {
		Code    uint64 `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Id         string `json:"id"`
			Detail     string `json:"detail"`
			ChainId    uint64 `json:"chainId"`
			Consensus  string `json:"consensus"`
			Difficulty string `json:"difficulty"`
			GasLimit   string `json:"gasLimit"`
			Alloc      Alloc  `json:"alloc"`
			Nodes      []Node `json:"nodes"`
		} `json:"data"`
	}{}
	err = json.Unmarshal(value, &response)
	if err != nil {
		return err
	}
	if response.Code != 200 {
		return errors.New(response.Message)
	}
	if response.Data.Consensus == "raft" {
		return bootRaft(out, consortium, privateKey, response.Data.Difficulty, response.Data.GasLimit, response.Data.ChainId, response.Data.Alloc, response.Data.Nodes)
	} else if response.Data.Consensus == "istanbul" {
		return bootIstanbul(out, consortium, privateKey, response.Data.Difficulty, response.Data.GasLimit, response.Data.ChainId, response.Data.Alloc, response.Data.Nodes)
	}
	return nil
}

func enodesRaft(nodes []Node) ([]byte, error) {
	enodes := []string{}
	for _, node := range nodes {
		enode := fmt.Sprintf(`enode://%s@%s:%d?discport=0&raftport=%d`, node.PublicKey, node.Host, node.Port, node.RaftPort)
		enodes = append(enodes, enode)
	}
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(enodes)
	if err != nil {
		return nil, err
	} else {
		return bf.Bytes(), nil
	}
}

func enodesIstanbul(nodes []Node) ([]byte, error) {
	enodes := []string{}
	for _, node := range nodes {
		enode := fmt.Sprintf(`enode://%s@%s:%d?discport=0`, node.PublicKey, node.Host, node.Port)
		enodes = append(enodes, enode)
	}
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(enodes)
	if err != nil {
		return nil, err
	} else {
		return bf.Bytes(), nil
	}
}

func bootRaft(out, consortium, privateKey, difficulty string, gasLimit string, chainId uint64, alloc Alloc, nodes []Node) error {
	genesis, err := raftGenesis(chainId, difficulty, gasLimit, alloc)
	if err != nil {
		return err
	}
	enodes, err := enodesRaft(nodes)
	if err != nil {
		return err
	}
	// node dir
	path := out + "/" + consortium + "-node"
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println(err)
	}

	// genesis.json
	err = common.CreateFile(path+"/genesis.json", genesis)
	if err != nil {
		return err
	}

	// static-nodes.json
	err = common.CreateFile(path+"/static-nodes.json", enodes)
	if err != nil {
		return err
	}

	// nodeKey
	err = common.CreateFile(path+"/nodeKey", []byte(common.HandlePrivateKey(privateKey)))
	if err != nil {
		return err
	}
	return nil
}

func bootIstanbul(out, consortium, privateKey, difficulty string, gasLimit string, chainId uint64, alloc Alloc, nodes []Node) error {
	genesis, err := istanbulGenesis(chainId, difficulty, gasLimit, alloc)
	if err != nil {
		return err
	}
	enodes, err := enodesIstanbul(nodes)
	if err != nil {
		return err
	}
	// node dir
	path := out + "/" + consortium + "-node"
	data := "/data"
	geth := "/geth"
	if err = os.MkdirAll(path+data+geth, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	// genesis.json
	err = common.CreateFile(path+"/genesis.json", genesis)
	if err != nil {
		return err
	}

	// static-nodes.json
	err = common.CreateFile(path+data+"/static-nodes.json", enodes)
	if err != nil {
		return err
	}
	// nodeKey
	err = common.CreateFile(path+data+geth+"/nodeKey", []byte(common.HandlePrivateKey(privateKey)))
	if err != nil {
		return err
	}
	return nil
}
