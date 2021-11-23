package server

import(
	pb "github.com/ap/DMP3/api"	
	"log"
	"net"
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

func(n *Node) UpdateBid(newBid int32) pb.BidReply_Outcome{
	if(n.HighestBid<newBid){
		n.HighestBid = newBid
		
		return pb.BidReply_SUCCESS
	} else if(n.HighestBid>=newBid){
		return pb.BidReply_FAIL
	}

	return pb.BidReply_EXCEPTION
}

func(n *Node) getResult() int32{
	return n.HighestBid
}