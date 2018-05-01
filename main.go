package main // import github.com/amckinney/protoc-enums

import (
	"bufio"
	"io/ioutil"
	"os"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	req := &plugin.CodeGeneratorRequest{}
	if err := req.Unmarshal(bytes); err != nil {
		panic(err)
	}

	res, err = generate(req)
	if err != nil {
		panic(err)
	}

	out, err := res.Marshal(make([]byte), true)
	if err != nil {
		panic(err)
	}

	_, err := os.Stdout.Write(out)
	if err != nil {
		panic(err)
	}
}

func generate(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	return &plugin.CodeGeneratorResponse{}, nil
}
