package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ap/DMP3/api"
	"github.com/ap/DMP3/internal/logging"
	"google.golang.org/grpc"
)

const (
	ip = ":5001"
)

type Node struct {
	HighestBid int32
	pb.UnimplementedAuctionServer
}

var (
	logger = logging.New()
)

func main() {
	node := &Node{
		HighestBid: 0,
	}

	node.StartServer()
}

func (n *Node) StartServer() {
	logger.IPrintf("Starting server\n")

	lis, err := net.Listen("tcp", ip)
	if err != nil {
		logger.EPrintf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuctionServer(s, n)

	logger.IPrintf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (n *Node) Bid(_ context.Context, req *pb.BidRequest) (*pb.BidReply, error) {
	logger.IPrintf("Retrieved bid request: %v, highest bid at this moment: %d\n", req, n.HighestBid)
	newBid := req.GetBid()

	if n.HighestBid < newBid {
		logger.IPrintf("Setting new highest value. Old: %d, new %d\n", n.HighestBid, newBid)
		n.HighestBid = newBid

		return &pb.BidReply{
			Outcome: pb.BidReply_SUCCESS,
		}, nil
	} else if n.HighestBid >= newBid {
		logger.IPrintf("New bid is below the highest bid. Highest: %d, new %d\n", n.HighestBid, newBid)
		return &pb.BidReply{
			Outcome: pb.BidReply_FAIL,
		}, nil
	}

	logger.EPrintf("Something bad happened. Highest: %d, new %d\n", n.HighestBid, newBid)
	return &pb.BidReply{
		Outcome: pb.BidReply_EXCEPTION,
	}, nil
}

func (n *Node) GetResult(_ context.Context, _ *pb.ResultRequest) (*pb.ResultReply, error) {
	logger.IPrintf("Retrieved get request. Highest bid at this moment: %d\n", n.HighestBid)
	return &pb.ResultReply{
		Result: n.HighestBid,
	}, nil
}
