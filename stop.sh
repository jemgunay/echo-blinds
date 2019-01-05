#!/bin/bash

fuser 3000/tcp > /dev/null 2>&1
if [[ $? != 0 ]]; then
    echo "> There is no echo-blinds service running on port 3000..."
    exit 1
fi

fuser -k 3000/tcp
