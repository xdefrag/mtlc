package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/stellar/go/clients/horizonclient"
)

var configPath = []string{
	"mtlc.toml",
	"$HOME/.config/mtlc/mtlc.toml",
	"$XDG_CONFIG_HOME/mtlc/mtlc.toml",
}

type Account struct {
	Address string `toml:"address"`
	Seed    string `toml:"seed"`
}

type Config struct {
	Testnet  bool      `toml:"testnet"`
	Accounts []Account `toml:"accounts"`
}

func (m *MTLC) Init() error {
	m.cfg = &Config{}

	for _, path := range configPath {
		if _, err := os.Stat(path); err == nil {
			if err := m.readConfig(path); err != nil {
				return err
			}
		}
	}

	if len(m.cfg.Accounts) == 0 {
		return fmt.Errorf("no accounts found in config")
	}

	m.cl = horizonclient.DefaultPublicNetClient
	if m.cfg.Testnet {
		m.cl = horizonclient.DefaultTestNetClient
	}

	return nil
}

func (m *MTLC) readConfig(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	if err := toml.NewDecoder(f).Decode(m.cfg); err != nil {
		return err
	}

	return nil
}
