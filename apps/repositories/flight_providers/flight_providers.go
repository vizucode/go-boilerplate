package flightproviders

import (
	"context"
	"fmt"
	"log"

	"github.com/vizucode/gokit/logger"
	pb "gitlab.com/tixia-backend/go-boilerplate/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type flightProvider struct {
	client pb.FlightResponseClient
}

func NewFlightProviders(host, port string) *flightProvider {
	addr := fmt.Sprintf("%s:%s", host, port)

	intercept := logger.NewInterceptor(addr)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(intercept.ChainUnaryClient(
			intercept.UnaryClientTracerInterceptor,
		)),
	}

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		logger.Log.Error(context.Background(), err)
		log.Fatal("grpc client could not connect to", addr, err)
	}

	OtaProvider := pb.NewFlightResponseClient(conn)

	return &flightProvider{
		client: OtaProvider,
	}
}
