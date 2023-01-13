#!/bin/sh

# Colors
Color_Off='\033[0m'       # Text Reset
Yellow='\033[0;33m'       # Yellow

# Process PID
PID_MAIN=`ps | grep '/opt/app/api/tmp/main' | grep -v grep | awk '{print $1}'`
PID_DLV=`ps | grep 'dlv --' | grep -v grep | awk '{print $1}'`

# Clear screen
clear

if [[ ! -z $PID_DLV ]]
then
    kill -9 $PID_DLV
fi


if [[ ! -z $PID_MAIN ]]
then
    printf "${Yellow}> DLV Service is listening... ${Color_Off} \n"

    dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient attach $PID_MAIN
fi
