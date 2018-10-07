package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vmihailenco/msgpack"
)

func msgpack2json(raw []byte) ([]byte, error) {
	var data interface{}
	err := msgpack.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}

func json2msgpack(raw []byte) ([]byte, error) {
	var data interface{}
	err := json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	return msgpack.Marshal(data)
}

func main() {
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "mpj: error:", err.Error())
		os.Exit(1)
	}

	var data []byte
	if json.Valid(raw) {
		data, err = json2msgpack(raw)
	} else {
		data, err = msgpack2json(raw)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "mpj: error:", err.Error())
		os.Exit(1)
	}

	os.Stdout.Write(data)
	fmt.Println()
}
