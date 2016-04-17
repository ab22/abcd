# Common functions to be used on tools.

printInfo() {
    echo -e "\e[0;33m + \e[1;32m$1 \e[0m"
}

printWarn() {
    echo -e "\e[0;33m + \e[1;33m$1 \e[0m"
}

printError() {
    echo -e "\e[0;33m + \e[1;31m$1 \e[0m"
    exit 1
}

promptMsg () {
    [ `echo $FLAGS | grep y` ] && return 0

    echo -en "\e[0;33m + \e[1;32m$1 \e[0m"
    read -p "[y/N]: " USER_RESPONSE

    if [ -z $USER_RESPONSE ]; then
        return 1
    fi

    if [[ $USER_RESPONSE =~ ^[Yy]$ ]]; then
        return 0
    else
        return 1
    fi
}
