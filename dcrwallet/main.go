package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/decred/dcrd/dcrutil"
	pb "github.com/decred/dcrwallet/rpc/walletrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var certificateFile = filepath.Join(dcrutil.AppDataDir("dcrwallet", false), "rpc.cert")

func main() {
	creds, err := credentials.NewClientTLSFromFile(certificateFile, "localhost")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := grpc.Dial("localhost:19558", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	c := pb.NewWalletLoaderServiceClient(conn)

	var seed [32]byte
	createWalletRequest := &pb.CreateWalletRequest{
		PrivatePassphrase: []byte("123"),
		Seed:              seed[:],
	}

	createWalletResponse, err := c.CreateWallet(context.Background(), createWalletRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Create/Import wallet: ", createWalletResponse)
}
