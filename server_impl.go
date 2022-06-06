package main

import (
	"MyID/MyID"
	"MyID/schemes"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Impl struct {
}

func (i *Impl) GenerateID(ctx context.Context, in *schemes.IDRequest) (*schemes.IDReply, error) {
	return &schemes.IDReply{
		Id:        MyID.IdService.Generate(),
		Timestamp: uint64(time.Now().UnixMicro()),
		ReplyId:   in.RequestId,
	}, nil
}

func Run(addr string) error {
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Printf("[MyID] %s", err.Error())
		return err
	}

	rpc := grpc.NewServer()

	schemes.RegisterMyIDServer(rpc, &Impl{})

	// reflection.Register(rpc)

	return rpc.Serve(listen)
}
