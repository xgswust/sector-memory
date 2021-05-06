package client

import (
	"context"
	"log"
	"os"

	pb "github.com/eben/sector-memory/proto"
	"google.golang.org/grpc"
)

// Client ..
type Client struct {
	DialAddr string
}

// NewClient ..
func NewClient() *Client {
	rpcAddr, ok := os.LookupEnv("STORAGE_LISTEN")
	if !ok {
		log.Println("NO STORAGE_LISTEN")
	}

	return &Client{
		DialAddr: rpcAddr,
	}
}

func (c *Client) connect() (pb.GrpcClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(c.DialAddr, grpc.WithInsecure()) //连接gRPC服务器
	if err != nil {
		return nil, nil, err
	}
	client := pb.NewGrpcClient(conn) //建立客户端
	return client, conn, nil
}

// GetSectorID ..
func (c *Client) ReportSectorID(ctx context.Context, sectorNum uint64, actorId uint64, primary bool, storageId string, sectorFileType int64) (uint64, error) {
	client, conn, err := c.connect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	req := new(pb.SectorMemRequest)

	req.SectorNumber = sectorNum
	req.ActorID = actorId
	req.Primary = primary
	req.StorageID = storageId
	req.SectorFileType = sectorFileType
	resp, err := client.DeclareSectorMemory(ctx, req) //调用方法
	if err != nil {
		return 0, err
	}
	return resp.Answer, nil
}
