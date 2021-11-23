package server

import (
	"context"
	"log"
	"net"

	pb "github.com/ap/DMP3/api"
	"google.golang.org/grpc"
)

const (
	ip = "127.0.0.1:5001"
)

type Node struct{
	HighestBid int32
	pb.UnimplementedAuctionServer
}

func main(){
	node := &Node{
		HighestBid: 0,
	}

	node.StartServer()
}

func(n *Node) StartServer(){
	lis, err := net.Listen("tcp", ip)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuctionServer(s, n)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func(n *Node) UpdateBid(_ context.Context, req *pb.BidRequest) (*pb.BidReply, error){
	newBid := req.GetBid()
	
	if(n.HighestBid<newBid){
		n.HighestBid = newBid
		
		return &pb.BidReply{
			Outcome: pb.BidReply_SUCCESS,
		},nil
	} else if(n.HighestBid>=newBid){
		return &pb.BidReply{
			Outcome: pb.BidReply_FAIL,
		},nil
	}

	return &pb.BidReply{
		Outcome: pb.BidReply_EXCEPTION,
	},nil
}

func(n *Node) getResult(context.Context, *pb.ResultRequest) (*pb.ResultReply, error){
	return &pb.ResultReply{
		Result: n.HighestBid,
	},nil
}