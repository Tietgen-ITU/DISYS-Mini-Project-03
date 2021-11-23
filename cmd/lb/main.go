package main

import (
	"context"
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
	index int
}

// main func
func main() {

	// Get list of replicas
	s := &LoadBalancer{}
	s.StartServer()
}

func (l *LoadBalancer) Bid(ctx context.Context, request *api.BidRequest) (*api.BidReply, error) {

	if l.isAuctionLive() {

		// TODO: Handle if all replicas are down...
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
	lis, err := net.Listen("tcp", "5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterAuctionServer(s, l)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
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
	// TODO: add IP

	url := strings.Split("", ":")[0] + "5000"
	log.Printf("Send res to: %s\n", url)

	conn, err := grpc.Dial(url, grpc.WithInsecure())
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
		log.Printf("Res errored: %v\n", err)
		return nil, err
	}

	return response, nil
}

// Send Res message
func (l *LoadBalancer) SendGetResult(endpoint string) (*api.ResultReply, error) {
	// TODO: add IP

	url := strings.Split("", ":")[0] + "5000"
	log.Printf("Send res to: %s\n", url)

	conn, err := grpc.Dial(url, grpc.WithInsecure())
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
		log.Printf("Res errored: %v\n", err)
		return nil, err
	}

	return response, nil
}

func (l *LoadBalancer) GetResult(context.Context, *api.ResultRequest) (*api.ResultReply, error) {
	index := l.index
	if l.index % 3 == 0 {
		l.index = 0
	} else {
		l.index += 1
	}

	return l.SendGetResult(l.replicaEndpoints[index])
}

// Handle incoming Req message
// func (n *Node) Req(ctx context.Context, in *pb.RequestMessage) (*pb.Empty, error) {
// 	callerIp := getClientIpAddress(ctx)

// 	log.Printf("Handling request from %s, current status: %d, local time: %d, in time: %d\n", callerIp, n.status, n.GetTs(), in.GetTime())

// 	status := n.GetStatus()
// 	if status == Status_HELD || (status == Status_WANTED && n.GetTs() < in.GetTime()) {
// 		log.Printf("Enque %s\n", callerIp)
// 		n.queue.Enqueue(callerIp)
// 		n.queue.Print()
// 	} else {
// 		n.SendRes(callerIp)
// 	}

// 	n.timestamp.Sync(in.GetTime())

// 	return &pb.Empty{}, nil
// }
