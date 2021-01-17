package main

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/prometheus/prometheus/prompb"
	"time"
)


func  unmarshal(reqBuf []byte) (prompb.WriteRequest, error) {
	var req prompb.WriteRequest
	if err := proto.Unmarshal(reqBuf, &req); err != nil {
		return req, err
	}

	return req, nil
}

func main() {
	fmt.Println("listening, ", time.Now())
}
