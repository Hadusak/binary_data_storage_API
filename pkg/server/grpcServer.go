package server

import (
	"github.com/Hadusak/binary_data_storage_API/pkg/utils"
	"log"
	"net"
)

func NewGRPCServer() {
	lis, err := net.Listen("tcp", utils.GetEnv("PROTO_PORT", ":31744"))
	if err != nil{
		log.Fatalf("failed to listen %v", err)
	}

	proto
}