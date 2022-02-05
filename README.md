# stockScanner
stockScanner application for tracking strong bullish and bearish stocks based on monthly, 21day, 6 month, bollinder bands

## stock variant
`go run main.go -t bullish -o file` should provide a csv file with all bullish stocks
`go run main.go -t bearish -o file` should provide a csv file with all bearish stocks

## simple option
`go run main.go -t bullish -o file -s` options s provides simple output instead of many technical values
if option s is not mentioned all technical values will be updated in csv sheet
 
 ## table view in command line
 `go run main.go -t bullish -o table -s` writes all data to a table format in command line

 ![Alt table_view](/assets/table_view.png?raw=true "table view" )

default option should provide bullish stocks, csv filename generated will be with time stamp for
example Bearish_20201014085412.csv if bearish else Bullish_20201014085412.csv

## mail option for sending in mail

`go run main.go -t bullish -o mail -s` option can be used for sending mail by configuring config/system.yaml
probably a cron job on raspberry pi or cloud server should be ideal usecase for this option. Configure all required
details make sure from email is enabled to send mails as low sensitive applications.

![Alt config](/assets/config.png?raw=true "config")