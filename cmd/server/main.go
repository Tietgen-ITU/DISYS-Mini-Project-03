package server

import(
	pb "github.com/ap/DMP3/api"	
)

type Node struct{
	HighestBid int32
}

func Main(){

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