# Echo Blinds

An Amazon Alexa skill server for controlling window blinds via GPIO pin controlled motors with a Raspberry Pi Zero.

Models for 3D printed parts can be found in the `models` directory.

Video demo: https://youtu.be/fgj6HONCQOs

<p align="center">
  <img src="/img/image_1.jpeg" width="275"/>
  <img src="/img/image_2.jpeg" width="275"/>
</p>

## Usage

```bash
# build for Linux & ARM
make build
# zip executables & required scripts
make package
# build, package and deploy to target host with install script
# note: must change "rpi_ip_address" value in Makefile first
make deploy
```
