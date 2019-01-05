#!/bin/bash

fuser 3000/tcp > /dev/null 2>&1
if [[ $? == 0 ]]; then
    echo "> There is already an echo-blinds service running on port 3000..."
    exit 1
fi

nohup ./echo-blinds &> echo-blinds.log &
