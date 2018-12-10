#!/bin/bash

rpi_ip_address="192.168.1.226"

# build echo-blinds & manual-blinds executables
cd ./cmd/echo-blinds && env GOOS=linux GOARCH=arm GOARM=5 go build
cd ../manual-blinds && env GOOS=linux GOARCH=arm GOARM=5 go build
# scp to raspberry pi
scp ../echo-blinds/echo-blinds manual-blinds pi@${rpi_ip_address}:/tmp/
