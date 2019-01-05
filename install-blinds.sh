#!/bin/bash

# move zip to target install directory
rm -f ./install-blinds.sh
mkdir -p ~/blinds && mv ./echo-blinds.zip ~/blinds

# unzip
cd ~/blinds
rm -rf *.sh *-blinds
unzip ./echo-blinds.zip

# clean up
rm ./echo-blinds.zip
