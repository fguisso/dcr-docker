package main

import (
	"context"
	"fmt"

	pb "github.com/decred/dcrwallet/rpc/walletrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var certificateFile = "rpc.cert"

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

	createWalletRequest := &pb.CreateWalletRequest{
		PrivatePassphrase: []byte("123"),
		Seed:              []byte("b280922d2cffda44648346412c5ec97f429938105003730414f10b01e1402eac"),
	}

	createWalletResponse, err := c.CreateWallet(context.Background(), createWalletRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Create/Import wallet: ", createWalletResponse)
	return
}
