package main

import (
	"fmt"
	"os"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/txnbuild"
	"github.com/urfave/cli/v2"
)

type MTLC struct {
	cl  *horizonclient.Client
	cfg *Config
}

const (
	flagTestnet = "testnet"
	flagBaseFee = "basefee"
	flagTimeout = "timeout"
)

func newApp(m *MTLC) *cli.App {
	return &cli.App{
		Name:  "mtlc",
		Usage: "Montelibero Command Line Interface",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  flagTestnet,
				Value: false,
				Usage: "use the test network",
			},
			&cli.Int64Flag{
				Name:  flagBaseFee,
				Value: txnbuild.MinBaseFee,
				Usage: "base fee for transactions in stroups",
			},
			&cli.Int64Flag{
				Name:  flagTimeout,
				Value: 0,
				Usage: "time bounds for transactions",
			},
		},
		Action:   m.Root,
		Commands: []*cli.Command{},
	}
}

func main() {
	m := &MTLC{}
	if err := m.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := newApp(m).Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
