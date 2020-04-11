#!/usr/bin/env bash



BUILD_BASE_URL="https://raw.githubusercontent.com/swayamraina/fcd/master/release/"

BUILD_VERSION="https://raw.githubusercontent.com/swayamraina/fcd/master/release/resources/latest_build"
DAEMON_CONFIG="https://raw.githubusercontent.com/swayamraina/fcd/master/launchd.config"
DAEMON_SCRIPT="https://raw.githubusercontent.com/swayamraina/fcd/master/fcd.sh"


echo "\n You are here :: `pwd` \n"
read -p "please enter the complete path to your fcd config : " path_to_fcd_config


curl -O ${BUILD_VERSION} 2> /dev/null
latest_build="$(cat latest_build)"

selection='y'
if [[ -d  ~/.fcd ]]; then
    if [[ -f ~/.fcd/latest_build ]]; then
        current_build="$(cat ~/.fcd/latest_build)"
        if [[ ${latest_build} > ${current_build} ]]; then
            echo "there is a new version available for fcd"
            echo "latest version is ${latest_build} while you are running ${current_build}"
            read -p "do you want to upgrade? Enter [y/n] : " selection
        fi
    fi
else
    echo "setting up the application runtime..."
    mkdir -p ~/.fcd
    mv latest_build ~/.fcd/
fi

cd ~/.fcd/
if [[ ${selection} -eq 'y' || ${selection} -eq 'Y' ]]; then
    rm -rf fcd
    echo "downloading latest version..."
    curl -O ${BUILD_BASE_URL}${latest_build} 2> /dev/null
    chmod +x ${latest_build}
    mv ${latest_build} fcd
fi

echo "setting up environment file..."
rm -rf ~/.fcd/fcd.env
echo "path=${path_to_fcd_config}" >> ~/.fcd/fcd.env

echo "fetching daemon config for restarts..."
curl -O ${DAEMON_CONFIG} 2> /dev/null
chmod +x launchd.config
sudo mv launchd.config /Library/LaunchAgents/dev.swayamraina.fcd.plist

echo "fetching daemon run script..."
curl -O ${DAEMON_SCRIPT} 2> /dev/null
sh fcd.sh