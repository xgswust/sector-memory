package server

import (
	"context"
	"log"
	"net"
	"os"
	"sync"

	pb "github.com/xgswust/sector-memory/proto"

	"google.golang.org/grpc"
)

// Service 定义我们的服务
type Service struct {
	SectorIDLk sync.RWMutex // 对应RPC调用GetSectorID，返回miner的sectorid
	SectorID   uint64
	SCFilePath string
}

func (s *Service)DeclareSectorMemory(ctx context.Context, req *pb.SectorMemRequest) (*pb.SectorMemResponse, error) {
	s.SectorIDLk.Lock()
	defer s.SectorIDLk.Unlock()
	log.Println("xjgw GetSectorID=",s.SectorID)
	log.Println("xjgw Greq.SectorFileType=",req.SectorFileType,"req.StorageID",req.StorageID,"req.ActorID",req.ActorID)
	return &pb.SectorMemResponse{Answer: req.SectorNumber}, nil
}

func (s *Service)StorageDropSectorMemory(ctx context.Context, req *pb.SectorMemRequest) (*pb.SectorMemResponse, error) {
	s.SectorIDLk.Lock()
	defer s.SectorIDLk.Unlock()
	log.Println("xjgw GetSectorID=",s.SectorID)
	log.Println("xjgw Greq.SectorFileType=",req.SectorFileType,"req.StorageID",req.StorageID,"req.ActorID",req.ActorID)
	return &pb.SectorMemResponse{Answer: req.SectorNumber}, nil
}



// Run ..
func Run(scFilePath string) {
	rpcAddr, ok := os.LookupEnv("STORAGE_LISTEN")
	if !ok {
		log.Println("NO STORAGE_LISTEN ENV")
		return
	}

	listener, err := net.Listen("tcp", rpcAddr) // 监听本地端口
	if err != nil {
		log.Println(err)
	}
	log.Println("grpc server Listing on", rpcAddr)

	grpcServer := grpc.NewServer() // 新建gRPC服务器实例
	server := &Service{            // 在gRPC服务器注册我们的服务
		SectorID:   1,
		SCFilePath: scFilePath,
	}
	pb.RegisterGrpcServer(grpcServer, server)

	if err = grpcServer.Serve(listener); err != nil { //用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
		log.Println(err)
	}
}
