package main // import github.com/amckinney/protoc-gen-enums

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	req, err := fileToGeneratorRequest(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to create code generator request: %v", err)
	}

	res, err := generate(req)
	if err != nil {
		return fmt.Errorf("failed to create code generator response: %v", err)
	}

	out, err := proto.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed ot marshal code generator response into bytes: %v", err)
	}

	_, err = os.Stdout.Write(out)
	if err != nil {
		return fmt.Errorf("failed to write code generator response bytes: %v", err)
	}
	return nil

}

func fileToGeneratorRequest(f *os.File) (*plugin.CodeGeneratorRequest, error) {
	r := bufio.NewReader(os.Stdin)
	in, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	req := &plugin.CodeGeneratorRequest{}
	return req, req.Unmarshal(in)
}

func generate(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	files := req.GetProtoFile()
	response := make([]*plugin.CodeGeneratorResponse_File, len(files))
	for i, f := range files {
		file, err := descriptorToGeneratorFile(f.GetName(), f)
		if err != nil {
			return nil, err
		}
		response[i] = file
	}
	return &plugin.CodeGeneratorResponse{
		File: response,
	}, nil
}

func descriptorToGeneratorFile(filename string, d *descriptor.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	var output string
	for _, e := range d.GetEnumType() {
		output += fmt.Sprintf("%s\n", e.GetName())
	}
	return &plugin.CodeGeneratorResponse_File{
		Name:    &filename,
		Content: &output,
	}, nil
}
