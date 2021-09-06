package main

import (
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	dpb "google.golang.org/protobuf/types/descriptorpb"
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f, err := os.Open("hello.pb")
	check(err)
	b, err := io.ReadAll(f)
	check(err)
	m := &dpb.FileDescriptorSet{}
	err = proto.Unmarshal(b, m)
	check(err)
	marshaler := protojson.MarshalOptions{Multiline: true}
	b, err = marshaler.Marshal(m)
	check(err)
	fmt.Println(string(b))
}
