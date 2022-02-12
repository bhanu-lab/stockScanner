#!/bin/bash

go build -o $PWD/bin/stockScanner
sudo cp $PWD/bin/stockScanner /usr/local/bin/

