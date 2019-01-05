#!/bin/bash

fuser 3000/tcp >/dev/null
if [[ $? != 0 ]]; then
    echo "> There is no echo-blinds service running on port 3000..."
    exit 1
fi

fuser -k 3000/tcp
