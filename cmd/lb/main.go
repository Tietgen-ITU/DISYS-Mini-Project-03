package main

import (
	"context"
	"flag"
	"net"
	"strings"
	"sync"
	goTime "time"

	"github.com/ap/DMP3/api"
	"github.com/ap/DMP3/internal/logging"
	"google.golang.org/grpc"
)

var (
	logger = logging.New()
)

type LoadBalancer struct {
	api.UnimplementedAuctionServer
	startTime        goTime.Time
	replicaEndpoints []string
	index            int
	roundRobinMutex  sync.Mutex
}

func main() {
	serverAddrStr := flag.String("serverAddr", "", "Server to connect to")
	servernames := strings.Split(*serverAddrStr, ",")

	// Get list of replicas
	s := &LoadBalancer{
		replicaEndpoints: servernames,
		index:            0,
		roundRobinMutex: sync.Mutex{},
	}
	s.StartServer()
}

func (l *LoadBalancer) Bid(ctx context.Context, request *api.BidRequest) (*api.BidReply, error) {

	if l.isAuctionLive() {

		for index, v := range l.replicaEndpoints {
			if len(v) > 0 {
				response, err := l.SendBid(v, request)

				if err != nil {
					logger.EPrintf("failed to listen: %v", err)
					defer l.declareReplicaDead(index)
				} else if response.Outcome == api.BidReply_EXCEPTION {
					defer l.declareReplicaDead(index)
				}
			}
		}

		return &api.BidReply{
			// SUCCESS
			Outcome: api.BidReply_SUCCESS,
		}, nil
	} else {
		return &api.BidReply{
			// FAIL
			Outcome: api.BidReply_FAIL,
		}, nil
	}
}

// server
func (l *LoadBalancer) StartServer() {
	lis, err := net.Listen("tcp", "5000")
	if err != nil {
		logger.EPrintf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterAuctionServer(s, l)
	logger.IPrintf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.EPrintf("failed to serve: %v", err)
	}

	l.startAuction()
}

func (l *LoadBalancer) startAuction() {
	l.startTime = goTime.Now()
}

func (l *LoadBalancer) isAuctionLive() bool {
	elapsed := goTime.Since(l.startTime)
	return elapsed.Minutes() >= 1
}

func (l *LoadBalancer) declareReplicaDead(replicaEnpointIndex int) {
	l.replicaEndpoints[replicaEnpointIndex] = ""
}

// Send Res message
func (l *LoadBalancer) SendBid(endpoint string, request *api.BidRequest) (*api.BidReply, error) {

	logger.IPrintf("Send res to: %s\n", endpoint)

	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	// client
	client := api.NewAuctionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), goTime.Second)
	defer cancel()

	response, err := client.Bid(ctx, request)
	if err != nil {
		logger.EPrintf("Res errored: %v\n", err)
		return nil, err
	}

	return response, nil
}

// Send Res message
func (l *LoadBalancer) SendGetResult(endpoint string) (*api.ResultReply, error) {

	logger.IPrintf("Send res to: %s\n", endpoint)

	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	// client
	client := api.NewAuctionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), goTime.Second)
	defer cancel()

	response, err := client.GetResult(ctx, &api.ResultRequest{})
	if err != nil {
		logger.EPrintf("Res errored: %v\n", err)
		return nil, err
	}

	return response, nil
}

/*
Gets the result by using a round robin
*/
func (l *LoadBalancer) GetResult(context.Context, *api.ResultRequest) (*api.ResultReply, error) {

	l.roundRobinMutex.Lock()
	defer l.roundRobinMutex.Unlock()

	index := l.index
	if l.index%3 == 0 {
		l.index = 0
	} else {
		l.index += 1
	}

	return l.SendGetResult(l.replicaEndpoints[index])
}
