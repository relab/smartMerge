package regserver

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/relab/smartMerge/proto"
	grpc "google.golang.org/grpc"
)

// For now, I don't imagine needing to start several servers.
// We therefore do this globally, instead of simply returning the grpcServer.
var grpcServer *grpc.Server
var mu sync.Mutex
var haveServer = false


func Start(port int) (*RegServer, error) {
	return StartInConf(port, nil, uint32(0))
}

func StartInConf(port int, init *pb.Blueprint, initC uint32) (*RegServer, error) {
	mu.Lock()
	defer mu.Unlock()
	if haveServer == true {
		log.Println("Abort start of grpc server, since old server exists.")
		return nil, errors.New("There already exists an old server.")
	}

	rs := NewRegServerWithCur(init, initC)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer = grpc.NewServer(opts...)
	pb.RegisterRegisterServer(grpcServer, rs)
	go grpcServer.Serve(lis)
	haveServer = true

	return rs, nil
}

func Stop() error {
	mu.Lock()
	defer mu.Unlock()

	if haveServer == false {
		log.Println("Tried to stop grpc-server, but no server was found.")
		return errors.New("No grpc server found.")
	}

	grpcServer.Stop()
	haveServer = false
	return nil
}

func StartTest(port int) (*grpc.Server, error) {
	rs := NewRegServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServ := grpc.NewServer(opts...)
	pb.RegisterRegisterServer(grpcServ, rs)
	go grpcServ.Serve(lis)
	haveServer = true

	return grpcServ, nil

}


////////////////// Advanced Server //////////////////////

func StartAdv(port int) (*RegServer, error) {
	return StartAdvInConf(port, nil, uint32(0))
}

func StartAdvInConf(port int, init *pb.Blueprint, initC uint32) (*RegServer, error) {
	mu.Lock()
	defer mu.Unlock()
	if haveServer == true {
		log.Println("Abort start of grpc server, since old server exists.")
		return nil, errors.New("There already exists an old server.")
	}

	rs := NewRegServerWithCur(init, initC)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer = grpc.NewServer(opts...)
	pb.RegisterAdvRegisterServer(grpcServer, rs)
	go grpcServer.Serve(lis)
	haveServer = true

	return rs, nil
}

func StartAdvTest(port int) (*grpc.Server, error) {
	rs := NewRegServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServ := grpc.NewServer(opts...)
	pb.RegisterAdvRegisterServer(grpcServ, rs)
	go grpcServ.Serve(lis)
	haveServer = true

	return grpcServ, nil

}