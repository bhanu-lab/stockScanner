# stockScanner

## stock variant
`go run main.go -t bullish` should provide a csv file with all bullish stocks
`go run main.go -t bearish` should provide a csv file with all bearish stocks

## simple option
`go run main.go -ts bullish` options s provides simple output instead of many technical values
if option s is not mentioned all technical values will be updated in csv sheet
 
 
default option should provide bullish stocks, csv filename generated will be with time stamp for
example Bearish_20201014085412.csv if bearish else Bullish_20201014085412.csv
