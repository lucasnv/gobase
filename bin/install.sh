#!/bin/sh
EXCLUDE_DIR=
# Colors
Color_Off='\033[0m'       # Text Reset
Yellow='\033[0;33m'       # Yellow
Green='\033[0;32m'        # Green

printf "\n"
printf "${Yellow}> Installing, wait a moment please... ${Color_Off} \n"
printf "\n"


# set the module URL
grep -rl --exclude-dir=.git \
         --exclude-dir=bin \
         --exclude-dir=.github \
         --exclude-dir=config \
         --exclude-dir=db \
         --exclude-dir=scripts \
         '<MODULE_URL_REPLACE>' . | xargs sed -i "s+<MODULE_URL_REPLACE>+$1+g"

# set the container name
grep -rl --exclude-dir=.git \
         --exclude-dir=bin \
         --exclude-dir=.github \
         --exclude-dir=config \
         --exclude-dir=db \
         --exclude-dir=scripts \
         '<CONTAINER_NAME_REPLACE>' . | xargs sed -i "s+<CONTAINER_NAME_REPLACE>+$2+g"         


printf "\n"
printf "${Green}> Installation complete. ${Color_Off} \n"
printf "\n"