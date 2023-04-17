package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	peerAddr := flag.String("peer-address", "", "peer address")
	flag.Parse()
	// ctx := context.Background()

	host, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))
	if err != nil {
		panic(err)
	}

	defer host.Close()

	fmt.Println("Addresses: ", host.Addrs())
	fmt.Println("ID: ", host.ID())

	if *peerAddr != "" {
		peerMA, err := multiaddr.NewMultiaddr(*peerAddr)
		if err != nil {
			panic(err)
		}

		peerAddrInfo, err := peer.AddrInfoFromP2pAddr(peerMA)
		if err != nil {
			panic(err)
		}

		if err := host.Connect(context.Background(), *peerAddrInfo); err != nil {
			panic(err)
		}

		fmt.Println("connected to", peerAddrInfo.String())

	}

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGKILL, syscall.SIGINT)
	<-sigCh
}