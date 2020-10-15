package server

import (
	context "context"
	"github.com/Hadusak/binary_data_storage_API/pkg/proto"
	"github.com/Hadusak/binary_data_storage_API/pkg/storage"
	"github.com/Hadusak/binary_data_storage_API/pkg/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

type DataServiceServer struct {
	Storage storage.Storage
}

func (d DataServiceServer) GetData(ctx context.Context, request *proto.GetDataRequest) (*proto.GetDataResponse, error) {
	data := d.Storage.Load(request.Key)
	return &proto.GetDataResponse{
		Ok:                   true,
		Message:              "",
		Data:                 &proto.Data{
			Key:                  request.Key,
			Data:                 data.Value,
			Timestamp:            data.Timestamp.Unix(),
		},
	}, nil
}

func (d DataServiceServer) SaveData(ctx context.Context, request *proto.SaveDataRequest) (*proto.SaveDataResponse, error) {
	key, data := utils.ProtoDataToInternalData(request.Data)

	d.Storage.Save(key, data)

	return &proto.SaveDataResponse{
		Ok:                   true,
		Message:              "",
	}, nil
}

func NewGRPCServer(storage storage.Storage) {
	lis, err := net.Listen("tcp", utils.GetEnv("PROTO_PORT", ":31744"))
	if err != nil{
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterDataServiceServer(grpcServer, DataServiceServer{Storage: storage})

	if err := grpcServer.Serve(lis); err != nil {
		//todo log and bye
	}
}

