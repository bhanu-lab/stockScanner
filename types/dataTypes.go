// types provide all data types used in stockScanner
package types

import (
	"gopkg.in/yaml.v2"
)

// used for determining whether a stock selected is Bearish or Bullish
const (
	DUMMY = iota
	BULLISH
	BEARISH
)

type Params struct {
	Key   string
	Value Value
}

type filter string

type Value struct {
	Action string
	Info   map[string]map[filter]string
}

type Config struct {
	Mail Mail `yaml:"mail"`
}

type Mail struct {
	From    string   `yaml:"from"`
	To      []string `yaml:"to"`
	Pass    string   `yaml:"pass"`
	Host    string   `yaml:"host"`
	Port    string   `yaml:"port"`
	Message []byte   `yaml:"-"`
}

func (c *Config) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

//InfoParams := make(map[string]map[filter]string)
