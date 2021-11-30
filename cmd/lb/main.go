package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	goTime "time"

	"github.com/ap/DMP3/api"
	"google.golang.org/grpc"
)

type LoadBalancer struct {
	api.UnimplementedAuctionServer
	startTime        goTime.Time
	replicaEndpoints []string
	index            int
}

func main() {
	serverAddrStr := flag.String("serverAddr", "abe123", "Server to connect to")
	flag.Parse()

	log.Printf("Input from flag: %s\n", *serverAddrStr)

	servernames := strings.Split(*serverAddrStr, ",")
	log.Printf("Replicas to forward reqeusts to: %v\n", servernames)

	// Get list of replicas
	s := &LoadBalancer{
		replicaEndpoints: servernames,
		index:            0,
	}
	s.StartServer()
}

func (l *LoadBalancer) Bid(ctx context.Context, request *api.BidRequest) (*api.BidReply, error) {

	if l.isAuctionLive() {

		for index, v := range l.replicaEndpoints {
			if len(v) > 0 {
				response, err := l.SendBid(v, request)

				if err != nil {
					log.Fatalf("failed to listen: %v", err)
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

	l.startAuction()
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterAuctionServer(s, l)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (l *LoadBalancer) startAuction() {
	l.startTime = goTime.Now()
	log.Printf("Starting auction at: %s\n", l.startTime)
}

func (l *LoadBalancer) isAuctionLive() bool {
	elapsed := goTime.Since(l.startTime)
	return elapsed.Minutes() < 1
}

func (l *LoadBalancer) declareReplicaDead(replicaEnpointIndex int) {
	l.replicaEndpoints[replicaEnpointIndex] = ""
}

// Send Res message
func (l *LoadBalancer) SendBid(endpoint string, request *api.BidRequest) (*api.BidReply, error) {

	if !l.isAuctionLive() {
		fmt.Println("Auction is finished! Denying bid request from %s", endpoint)

		return &api.BidReply{
			Outcome: api.BidReply_FAIL,
		}, nil
	}

	log.Printf("Send bid %v to: %s\n", request.Bid, endpoint)

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
		log.Printf("Bid errored: %v\n", err)
		return nil, err
	}

	return response, nil
}

// Send Res message
func (l *LoadBalancer) SendGetResult(endpoint string) (*api.ResultReply, error) {

	log.Printf("Send GetResult to: %s\n", endpoint)

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
		log.Printf("GetResult errored: %v\n", err)
		return nil, err
	}

	return response, nil
}

func (l *LoadBalancer) GetResult(context.Context, *api.ResultRequest) (*api.ResultReply, error) {
	index := l.index
	if l.index%3 == 0 {
		l.index = 0
	} else {
		l.index += 1
	}

	return l.SendGetResult(l.replicaEndpoints[index])
}
