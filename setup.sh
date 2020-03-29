#!/usr/bin/env bash


echo ""
echo "You are here :: `pwd`"
echo ""
echo "please enter the complete path to your fcd config : "

# input path to fcd_config
read path_to_fcd_config

# create env file
echo ${path_to_fcd_config} >> ~/.fcd_env