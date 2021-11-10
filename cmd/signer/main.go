package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/POKTBridge/signer/config"
	"github.com/POKTBridge/signer/service"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	grpcsign "github.com/poktbridge/protocols/sign"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.AutoLoad()
	if err != nil {
		log.Fatalln(err)
	}

	decETHPk, err := base64.StdEncoding.DecodeString(cfg.EthPk)
	if err != nil {
		log.Fatalln(err)
	}

	ethPK := service.NewPK(decETHPk)

	decPoktPk, err := base64.StdEncoding.DecodeString(cfg.PoktPk)
	if err != nil {
		log.Fatalln(err)
	}

	poktPK := service.NewPK(decPoktPk)

	chainService := service.NewSign(ethPK, poktPK)

	grpcMetrics := grpc_prometheus.DefaultServerMetrics

	grpcWrapper := service.NewServerWrapper(chainService, cfg.EthPass, cfg.PoktPass)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	lis, listErr := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if listErr != nil {
		wg.Done()
		log.Fatalln(err)
	}
	log.Printf("grpc server listening at %v\n", lis.Addr())

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)

	grpcsign.RegisterSignServer(s, grpcWrapper)
	grpc_prometheus.Register(s)
	s.Serve(lis)
}
