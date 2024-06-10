package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func (m *MTLC) Root(c *cli.Context) error {
	fmt.Fprintln(c.App.Writer, "Accounts:")

	for _, account := range m.cfg.Accounts {
		fmt.Fprintf(c.App.Writer, "\t%s\n", account.Address)
	}

	return nil
}
