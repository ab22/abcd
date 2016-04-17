#!/bin/bash

# Utility script used to copy the necessary files to the Dropbox folder
# before deploying them to Heroku.
#
# Before running this command it is recommended that the Dropbox
# client is paused on the local machine to avoid unnecessary file
# uploads to Dropbox before finishing copying/removing all files.
#
#

# Exit if any of the commands returns an error.
set -e

SRC_BACKEND_FOLDER=$GOPATH"/src/github.com/ab22/abcd/"
SRC_FRONTEND_FOLDER=$SRC_BACKEND_FOLDER"/frontend/admin/"

DEST_BACKEND_FOLDER=$HOME"/Dropbox/Aplicaciones/Heroku/instituto-abcd/"
DEST_FRONTEND_FOLDER=$DEST_BACKEND_FOLDER"/frontend/admin/"

BACKEND_FILE_LIST=(
    "main.go"
    "Procfile"
    "server.go"
)

BACKEND_FOLDER_LIST=(
    "config/"
    "Godeps/"
    "handlers/"
    "httputils/"
    "models/"
    "routes/"
    "services/"
    "vendor/"
)

FRONTEND_FOLDER_LIST=(
    "dist/"
)

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

deployFrontend() {
    printInfo "Deploying frontend..."

    mkdir -p $DEST_FRONTEND_FOLDER

    for folder in "${FRONTEND_FOLDER_LIST[@]}"
    do
        cp -vr "$SRC_FRONTEND_FOLDER${folder}" ${DEST_FRONTEND_FOLDER}
    done

    printInfo "Frontend deployed!"
}

deployBackend(){
    printInfo "Deploying backend..."
    printInfo "Copying folders..."

    for folder in "${BACKEND_FOLDER_LIST[@]}"
    do
        cp -vr "$SRC_BACKEND_FOLDER${folder}" $DEST_BACKEND_FOLDER
    done

    printInfo "Copying files..."

    for file in "${BACKEND_FILE_LIST[@]}"
    do
        cp -v "$SRC_BACKEND_FOLDER${file}" ${DEST_BACKEND_FOLDER}
    done

    printInfo "Backend deployed!"
}

removeExistingFiles() {
    if [ $DEST_BACKEND_FOLDER = "/" ]; then
        printError "The destination deployment folder cannot be '/'"
    fi

    printInfo "Removing existing files from '${DEST_BACKEND_FOLDER}'..."
    rm -vrf "$DEST_BACKEND_FOLDER"/*
    printInfo "All files removed!"
}
deploy() {
    removeExistingFiles
    deployBackend
    deployFrontend
}

main() {
    printInfo "=================================================="
    printInfo "             Deployment tool v.0.0.3"
    printInfo "=================================================="
    printInfo "      Backend source folder: $SRC_BACKEND_FOLDER"
    printInfo " Backend destination folder: $DEST_BACKEND_FOLDER"
    printInfo "     Frontend source folder: $SRC_BACKEND_FOLDER"
    printInfo "Frontend destination folder: $DEST_FRONTEND_FOLDER"
    printInfo "=================================================="
    printWarn "This utility will delete all contents from the '$DEST_BACKEND_FOLDER' and"
    printWarn "will replace them with the content of '$SRC_BACKEND_FOLDER'!"
    printWarn "Reminder: Make sure to run 'grunt build' on the frontend folder!"

    if promptMsg "Are you sure you want to continue?"; then
        printInfo "Starting deployment..."
        deploy
        printInfo "Deployment done!"
    else
        printError "Deployment aborted!"
    fi
}

main
