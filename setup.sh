#!/usr/bin/env bash


echo ""
echo "You are here :: `pwd`"
echo ""
echo "please enter the complete path to your fcd config : "

# input path to fcd_config
read path_to_fcd_config

# create env file
echo ${path_to_fcd_config} >> ~/.fcd_env

# create directory
mkdir -p ~/.fcd

# get latest build
latest_build="$(curl https://raw.githubusercontent.com/swayamraina/fcd/master/release/latest_build)"

# download latest build
curl -O https://raw.githubusercontent.com/swayamraina/fcd/master/release/${latest_build}

# mark as executable
chmod +x ${latest_build}

# move to the new directory
mv ${latest_build} ~/.fcd/

# move pointer to executable location
cd ~/.fcd/

# run fcd
./${latest_build} &
