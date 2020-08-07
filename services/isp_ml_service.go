package services

import (
	"context"
	"google.golang.org/grpc"
	pb "gin-restful-best-practice/protos/ml_service"
	"log"
	"time"
)

var conn *grpc.ClientConn

func DownloadModel(name string) ([]byte, error) {
	client := pb.NewMachineLearningServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.DownloadModel(ctx, &pb.DownloadModelRequest{Name: name})
	if resp == nil {
		return nil, err
	}
	return resp.Model, err
}

func connect() (err error) {
	conn, err = grpc.Dial("localhost:50051",
		grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(500*time.Millisecond))
	return err
}

func init() {
	log.Println("Establishing isp_ml_service connection...")
	if err := connect(); err != nil {
		log.Printf("isp_ml_service: connecting failed: %v\n", err)
	}
}
