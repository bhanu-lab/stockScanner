# stockScanner

## stock variant
`go run main.go -t bullish -o file` should provide a csv file with all bullish stocks
`go run main.go -t bearish -o file` should provide a csv file with all bearish stocks

## simple option
`go run main.go -t bullish -o file -s` options s provides simple output instead of many technical values
if option s is not mentioned all technical values will be updated in csv sheet
 
 ## table view in command line
 `go run main.go -t bullish -o table -s` writes all data to a table format in command line

default option should provide bullish stocks, csv filename generated will be with time stamp for
example Bearish_20201014085412.csv if bearish else Bullish_20201014085412.csv
