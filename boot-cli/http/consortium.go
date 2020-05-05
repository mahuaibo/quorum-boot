package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mahuaibo/quorum-boot/boot-cli/common"
)

type Alloc map[string]Account

type Account struct {
	Balance string `json:"balance"`
}

type Node struct {
	PublicKey string `json:"publicKey"`
	Host      string `json:"host"`
	Port      uint64 `json:"port"`
	RaftPort  uint64 `json:"raftport"`
}

type Istanbul struct {
	Epoch          uint64 `json:"epoch"`
	Policy         uint64 `json:"policy"`
	Ceil2Nby3Block uint64 `json:"ceil2Nby3Block"`
}

type CodeSizeConfig struct {
	Block uint64 `json:"block"`
	Size  uint64 `json:"size"`
}

func GetConsortium(serverUrl, consortium string) ([]byte, []byte, error) {
	value, err := common.HttpGet(serverUrl, "/consortiums/"+consortium+"/boot")
	if err != nil {
		return nil, nil, err
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
		return nil, nil, err
	}
	if response.Code != 200 {
		return nil, nil, errors.New(response.Message)
	}
	var genesis []byte
	if response.Data.Consensus == "raft" {
		genesis, err = raftGenesis(response.Data.ChainId, response.Data.Difficulty, response.Data.GasLimit, response.Data.Alloc)
	} else if response.Data.Consensus == "istanbul" {
		genesis, err = istanbulGenesis(response.Data.ChainId, response.Data.Difficulty, response.Data.GasLimit, response.Data.Alloc)
	}
	if err != nil {
		return nil, nil, err
	}

	enodes := []string{}
	for _, node := range response.Data.Nodes {
		enode := fmt.Sprintf(`enode://%s@%s:%d?discport=0&raftport=%d`, node.PublicKey, node.Host, node.Port, node.RaftPort)
		enodes = append(enodes, enode)
	}
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(enodes)
	if err != nil {
		return nil, nil, err
	}
	return genesis, bf.Bytes(), err
}

// istanbul
type GenesisIstanbul struct {
	Alloc      Alloc          `json:"alloc"`    // 0x
	Coinbase   string         `json:"coinbase"` // 0x
	Config     ConfigIstanbul `json:"config"`
	Difficulty string         `json:"difficulty"` // 0x
	ExtraData  string         `json:"extraData"`  // 0x
	GasLimit   string         `json:"gasLimit"`   // 0x
	Mixhash    string         `json:"mixhash"`    // 0x
	Nonce      string         `json:"nonce"`      // 0x
	ParentHash string         `json:"parentHash"` // 0x
	Timestamp  string         `json:"timestamp"`  // 0x
	Number     string         `json:"number"`
	GasUsed    string         `json:"gasUsed"`
}
type ConfigIstanbul struct {
	HomesteadBlock      uint64   `json:"homesteadBlock"`
	ByzantiumBlock      uint64   `json:"byzantiumBlock"`
	ConstantinopleBlock uint64   `json:"constantinopleBlock"`
	ChainId             uint64   `json:"chainId"`
	Eip150Block         uint64   `json:"eip150Block"`
	Eip155Block         uint64   `json:"eip155Block"`
	Eip150Hash          string   `json:"eip150Hash"`
	Eip158Block         uint64   `json:"eip158Block"`
	IsQuorum            bool     `json:"isQuorum"`
	Istanbul            Istanbul `json:"istanbul"`
	TxnSizeLimit        uint64   `json:"txnSizeLimit"`
	MaxCodeSize         uint64   `json:"maxCodeSize"`
}

func istanbulGenesis(chainId uint64, difficulty string, gasLimit string, alloc Alloc) ([]byte, error) {
	genesis := &GenesisIstanbul{
		Alloc:    alloc,
		Coinbase: "0x0000000000000000000000000000000000000000",
		Config: ConfigIstanbul{
			HomesteadBlock:      0,
			ByzantiumBlock:      0,
			ConstantinopleBlock: 0,
			ChainId:             chainId,
			Eip150Block:         0,
			Eip155Block:         0,
			Eip150Hash:          "0x0000000000000000000000000000000000000000000000000000000000000000",
			Eip158Block:         0,
			Istanbul: Istanbul{
				Epoch:          30000,
				Policy:         0,
				Ceil2Nby3Block: 0,
			},
			IsQuorum:     true,
			TxnSizeLimit: 64,
			MaxCodeSize:  0,
		},
		Difficulty: difficulty,
		ExtraData:  "0x0000000000000000000000000000000000000000000000000000000000000000",
		GasLimit:   gasLimit,
		Mixhash:    "0x0000000000000000000000000000000000000000000000000000000000000000",
		Nonce:      "0x0",
		ParentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
		Timestamp:  "0x00",
		Number:     "0x0",
		GasUsed:    "0x0",
	}
	jsonValue, err := json.Marshal(genesis)
	if err != nil {
		return nil, err
	}
	var genesisBuffer bytes.Buffer
	err = json.Indent(&genesisBuffer, jsonValue, "", "\t")
	return genesisBuffer.Bytes(), err
}

// raft
type GenesisRaft struct {
	Alloc      Alloc      `json:"alloc"`    // 0x
	Coinbase   string     `json:"coinbase"` // 0x
	Config     ConfigRaft `json:"config"`
	Difficulty string     `json:"difficulty"` // 0x
	ExtraData  string     `json:"extraData"`  // 0x
	GasLimit   string     `json:"gasLimit"`   // 0x
	Mixhash    string     `json:"mixhash"`    // 0x
	Nonce      string     `json:"nonce"`      // 0x
	ParentHash string     `json:"parentHash"` // 0x
	Timestamp  string     `json:"timestamp"`  // 0x
}

type ConfigRaft struct {
	HomesteadBlock      uint64           `json:"homesteadBlock"`
	ByzantiumBlock      uint64           `json:"byzantiumBlock"`
	ConstantinopleBlock uint64           `json:"constantinopleBlock"`
	PetersburgBlock     uint64           `json:"petersburgBlock"`
	IstanbulBlock       uint64           `json:"istanbulBlock"`
	ChainId             uint64           `json:"chainId"`
	Eip150Block         uint64           `json:"eip150Block"`
	Eip155Block         uint64           `json:"eip155Block"`
	Eip150Hash          string           `json:"eip150Hash"`
	Eip158Block         uint64           `json:"eip158Block"`
	MaxCodeSizeConfig   []CodeSizeConfig `json:"maxCodeSizeConfig"`
	IsQuorum            bool             `json:"isQuorum"`
}

func raftGenesis(chainId uint64, difficulty string, gasLimit string, alloc Alloc) ([]byte, error) {
	genesis := &GenesisRaft{
		Alloc:    alloc,
		Coinbase: "0x0000000000000000000000000000000000000000",
		Config: ConfigRaft{
			HomesteadBlock:      0,
			ByzantiumBlock:      0,
			ConstantinopleBlock: 0,
			PetersburgBlock:     0,
			IstanbulBlock:       0,
			ChainId:             chainId,
			Eip150Block:         0,
			Eip155Block:         0,
			Eip150Hash:          "0x0000000000000000000000000000000000000000000000000000000000000000",
			Eip158Block:         0,
			MaxCodeSizeConfig: []CodeSizeConfig{
				{
					Block: 0,
					Size:  35,
				},
			},
			IsQuorum: true,
		},
		Difficulty: difficulty,
		ExtraData:  "0x0000000000000000000000000000000000000000000000000000000000000000",
		GasLimit:   gasLimit,
		Mixhash:    "0x0000000000000000000000000000000000000000000000000000000000000000",
		Nonce:      "0x0",
		ParentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
		Timestamp:  "0x00",
	}
	jsonValue, err := json.Marshal(genesis)
	if err != nil {
		return nil, err
	}
	var genesisBuffer bytes.Buffer
	err = json.Indent(&genesisBuffer, jsonValue, "", "\t")
	return genesisBuffer.Bytes(), err
}
