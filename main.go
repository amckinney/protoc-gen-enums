package main // import github.com/amckinney/protoc-gen-enums

import (
	"bufio"
	"io/ioutil"
	"os"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	in, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	req := &plugin.CodeGeneratorRequest{}
	if err := req.Unmarshal(in); err != nil {
		panic(err)
	}

	res, err := generate(req)
	if err != nil {
		panic(err)
	}

	bytes := []byte{}
	out, err := res.Marshal(bytes, true)
	if err != nil {
		panic(err)
	}

	_, err = os.Stdout.Write(out)
	if err != nil {
		panic(err)
	}
}

func generate(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	return &plugin.CodeGeneratorResponse{}, nil
}
