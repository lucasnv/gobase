#!/bin/sh
EXCLUDE_DIR=
# Colors
Color_Off='\033[0m'       # Text Reset
Yellow='\033[0;33m'       # Yellow

printf "\n"
printf "${Yellow}> Installing proyect ... ${Color_Off} \n"
printf "\n"


grep -rl --exclude-dir=.git --exclude-dir=bin --exclude-dir=.github --exclude-dir=cmd --exclude-dir=config --exclude-dir=db --exclude-dir=scripts "[MODULE_URL]" .


# find . -type f -exec sed -i 's+[MODULE_URL]+PRUEBAAAAAAAA+g' {} +

#grep -rl '[MODULE_URL]' bin/test/ | xargs sed -i "s+[MODULE_URL]+$1+g"
# grep -rl [CONTAINER_NAME] test/ | xargs sed -i "s+[CONTAINER_NAME]+$2+g"
