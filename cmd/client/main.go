package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"time"

	pb "github.com/ap/DMP3/api"
	"github.com/ap/DMP3/internal/logging"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("serverAddr", "localhost:5001", "Server to connect to")
	random     = flag.Bool("random", false, "Randomly send data")
	logger     = logging.New()
)

func main() {
	flag.Parse()

	logger.IPrintf("Dialing %s\n", *serverAddr)

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.EPrintf("Could not connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewAuctionClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if !*random {
		autoAuction(c, ctx)
	} else {
		auction(c, ctx)
	}
}

func autoAuction(c pb.AuctionClient, ctx context.Context) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		result, err := result(c, ctx)
		if err == nil {
			break
		}

		result += r.Int31n(10)
		err = bid(result, c, ctx)
		if err == nil {
			break
		}
	}
}

func auction(c pb.AuctionClient, ctx context.Context) {
	for {
		var choice, toBid int32
		logger.IPrintf("Do you want to get the result (1) or bid (2): ")
		if _, err := fmt.Scanf("%d\n", &choice); err != nil {
			logger.EPrintf("Invalid input, dying: %v\n", err)
		}

		switch choice {
		case 1:
			_, err := result(c, ctx)
			if err != nil {
				return
			}
			break
		case 2:
			logger.IPrintf("How much do you want to bet?\n> ")

			if _, err := fmt.Scanf("%d\n", &toBid); err != nil {
				logger.EPrintf("Invalid input, dying: %v\n", err)
			}

			err := bid(toBid, c, ctx)
			if err != nil {
				return
			}

			break
		default:
			logger.EPrintf("Invalid choice\n")
			return
		}

	}
}

func bid(amount int32, c pb.AuctionClient, ctx context.Context) error {
	logger.IPrintf("Bidding %d\n", amount)

	reply, err := c.Bid(ctx, &pb.BidRequest{
		Bid: amount,
	})

	if err != nil {
		logger.EPrintf("Failed to bid: %v\n", err)
		return err
	} else if reply.Outcome == pb.BidReply_FAIL || reply.Outcome == pb.BidReply_EXCEPTION {
		msg := fmt.Sprintf("Bid failed with outcome: %s", outcomeToString(reply.Outcome))
		logger.EPrintf("%s\n", msg)

		return errors.New(msg)
	}

	logger.IPrintf("Successfully bid %d\n", amount)

	return nil
}

func result(c pb.AuctionClient, ctx context.Context) (int32, error) {
	logger.IPrintf("Retrieving result\n")

	reply, err := c.GetResult(ctx, &pb.ResultRequest{})
	if err != nil {
		logger.EPrintf("Failed to retrieve result: %v\n", err)
		return 0, nil
	}

	logger.IPrintf("Retrieved result: %d\n", reply.Result)
	return reply.Result, nil
}

func outcomeToString(outcome pb.BidReply_Outcome) string {
	switch outcome {
	case pb.BidReply_SUCCESS:
		return "SUCCESS"
	case pb.BidReply_EXCEPTION:
		return "EXCEPTION"
	case pb.BidReply_FAIL:
		return "FAIL"
	default:
		return ""
	}
}
