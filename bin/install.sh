#!/bin/sh

# Colors
Color_Off='\033[0m'       # Text Reset
Yellow='\033[0;33m'       # Yellow

printf "\n"
printf "${Yellow}> Installing proyect ... ${Color_Off} \n"
printf "\n"

grep -rl '[MODULE_URL]' bin/test/ | xargs sed -i "s+[MODULE_URL]+$1+g"
# grep -rl [CONTAINER_NAME] test/ | xargs sed -i "s+[CONTAINER_NAME]+$2+g"
