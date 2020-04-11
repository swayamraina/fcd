#!/usr/bin/env bash



setup=false

if [[ -d ~/.fcd/ ]]; then
    setup=true
    cd ~/.fcd/
    if [[ -f ~/.fcd/fcd ]]; then
        setup=true
        pid=`ps aux | grep 'fcd' | grep -v grep | grep -v fcd.sh | awk '{print $2}'`
        if [[ "${pid}" -ne "" ]]; then
            echo "stopping existing process..."
            kill -9 ${pid}
        fi
        echo "running fcd..."
        ./fcd &
    else
        setup=false
        echo "executable does not exists"
    fi
fi

if [[ "${setup}" == false ]]; then
    echo "flash-card daemon not setup. Please check https://github.com/swayamraina/fcd"
fi