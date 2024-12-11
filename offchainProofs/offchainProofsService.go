package offchainProofs

import (
	"context"
	"fmt"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"github.com/0xPolygonHermez/zkevm-aggregator/log"
	"github.com/0xPolygonHermez/zkevm-aggregator/offchainProofs/protocols/offchain-proofs"
)

// IOffchainProofsService is an offchain proofs service interface
type IOffchainProofsService interface {
	VerifyBatchesSimulate(ctx context.Context, req *proofs_pb.VerifyBatchesRequest) error
	// Close is used to close the service
	Close()
}

// offchainProofsService is an IOffchainProofsService implementation
type offchainProofsService struct {
	address string
	conn    *grpc.ClientConn
	service proofs_pb.OffchainProofsServiceClient
}

// NewOffchainProofsService is used to get new IOffchainProofsService implementation instance
func NewOffchainProofsService(config Config) (IOffchainProofsService, error) {
	s := &offchainProofsService{address: config.Address}

	if err := s.init(); err != nil {
		return nil, err
	}

	return s, nil
}

// init is used to init the offchainProofsService
func (s *offchainProofsService) init() (err error) {
	kacp := keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	s.conn, err = grpc.NewClient(
		s.address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp),
		grpc.WithDefaultCallOptions(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	)
	if err != nil {
		return fmt.Errorf("err dial grpc %s: %w", s.address, err)
	}

	s.service = proofs_pb.NewOffchainProofsServiceClient(s.conn)

	log.Info("offchainProofsService:", "init on", s.address)

	return nil
}

func (s *offchainProofsService) VerifyBatchesSimulate(ctx context.Context, req *proofs_pb.VerifyBatchesRequest) error {
	log.Info("offchainProofsService: sending verify batch request")

	if _, err := s.service.VerifyBatchesSimulate(ctx, req); err != nil {
		return err
	}

	return nil
}

func (s *offchainProofsService) Close() {
	if s.conn != nil {
		if err := s.conn.Close(); err != nil {
			log.Warnf("err close grpc connection: %s", err.Error())
		}
	}
}
