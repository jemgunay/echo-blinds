# IP address of target raspberry pi for deployment
rpi_ip_address="<host_ip_address>"

clean:
	rm -rf ./build

build: clean
	mkdir build
	chmod +x install-blinds.sh

	# build echo-blinds & manual-blinds ARM executables
	cd ./cmd/echo-blinds && rm -f echo-blinds && env GOOS=linux GOARCH=arm GOARM=5 go build
	cd ./cmd/manual-blinds && rm -f manual-blinds && env GOOS=linux GOARCH=arm GOARM=5 go build
	cp -f ./cmd/*-blinds/*-blinds ./build

	# copy startup script
	cp -f start.sh stop.sh ./build
	chmod +x ./build/start.sh ./build/stop.sh

package:
	# package into zip
	rm -f ./build/echo-blinds.zip
	cd ./build && zip -r ./echo-blinds.zip ./*

upload:
	# scp zip & install script to raspberry pi
	scp ./build/echo-blinds.zip ./install-blinds.sh pi@${rpi_ip_address}:/tmp/

deploy: build package upload
