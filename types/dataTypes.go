// types provide all data types used in stockScanner
package types

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

//InfoParams := make(map[string]map[filter]string)
