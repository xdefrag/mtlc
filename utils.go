package main

import (
	"github.com/samber/lo"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
	"github.com/urfave/cli/v2"
)

func (m *MTLC) submit(c *cli.Context, ops []txnbuild.Operation) error {
	var (
		baseFee    = c.Int64(flagBaseFee)
		timeout    = c.Int64(flagTimeout)
		testnet    = c.Bool(flagTestnet)
		passphrase = lo.Ternary(testnet, network.TestNetworkPassphrase, network.PublicNetworkPassphrase)
	)

	tx, err := txnbuild.NewTransaction(txnbuild.TransactionParams{
		IncrementSequenceNum: true,
		BaseFee:              baseFee,
		Preconditions: txnbuild.Preconditions{
			TimeBounds: txnbuild.NewTimeout(timeout),
		},
		Operations: ops,
	})
	if err != nil {
		return err
	}

	signee, err := keypair.ParseFull(m.cfg.Accounts[0].Seed)
	if err != nil {
		return err
	}

	tx, err = tx.Sign(passphrase, signee)
	if err != nil {
		return err
	}

	_, err = m.cl.SubmitTransaction(tx)

	return err
}
