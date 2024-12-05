package main

import (
	"bytes"
	"testing"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestRoot(t *testing.T) {
	t.Parallel()

	m := newMTLCTest(t)
	a, out := newAppTest(t, m)

	acc := m.cfg.Accounts[0]

	err := a.Run([]string{"mtlc"})
	require.NoError(t, err)

	require.Contains(t, out.String(), acc.Address)

	t.Logf(out.String())
}

func newMTLCTest(t *testing.T) *MTLC {
	t.Helper()

	m := &MTLC{
		cl:  horizonclient.DefaultTestNetClient,
		cfg: &Config{},
	}

	acc := keypair.MustRandom()

	_, err := m.cl.Fund(acc.Address())
	require.NoError(t, err)

	m.cfg.Accounts = append(m.cfg.Accounts, Account{
		Address: acc.Address(),
		Seed:    acc.Seed(),
	})

	return m
}

func newAppTest(t *testing.T, m *MTLC) (*cli.App, *bytes.Buffer) {
	t.Helper()

	a := newApp(m)
	out := bytes.NewBuffer([]byte(""))
	a.Writer = out

	return a, out
}
