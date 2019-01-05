# IP address of target blinds raspberry pi
rpi_ip_address="192.168.1.226"

clean:
	rm -rf ./build

build: clean
	mkdir build

	# build echo-blinds & manual-blinds ARM executables
	cd ./cmd/echo-blinds && rm -f echo-blinds && env GOOS=linux GOARCH=arm GOARM=5 go build
	cd ./cmd/manual-blinds && rm -f manual-blinds && env GOOS=linux GOARCH=arm GOARM=5 go build
	cp -f ./cmd/*-blinds/*-blinds ./build

	# copy startup script
	cp -f start.sh stop.sh ./build
	chmod +x ./build/start.sh ./build/stop.sh

upload:
	# scp to raspberry pi
	scp ./build/*-blinds start.sh pi@${rpi_ip_address}:/tmp/

publish: build upload
