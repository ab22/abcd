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

# Determine where the current script is stored at so that we
# can source other scripts from it.
SCRIPT_PATH="$(dirname `which $0`)/"

# Load common functions.
source $SCRIPT_PATH"common.sh"

SRC_BACKEND_FOLDER=$GOPATH"/src/github.com/ab22/abcd/"
SRC_FRONTEND_FOLDER=$SRC_BACKEND_FOLDER"frontend/admin/"

DEST_BACKEND_FOLDER=$HOME"/Dropbox/Aplicaciones/Heroku/instituto-abcd/"
DEST_FRONTEND_FOLDER=$DEST_BACKEND_FOLDER"frontend/admin/"

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

buildFrontend() {
    printInfo "Running 'grunt build' on '$SRC_FRONTEND_FOLDER'..."

    cd $SRC_FRONTEND_FOLDER
    grunt build

    printInfo "Frontend successfully built!"
}

main() {
    printInfo "=================================================="
    printInfo "             Deployment tool v.0.2.1"
    printInfo "=================================================="
    printInfo "       Script location path: $SCRIPT_PATH"
    printInfo "      Backend source folder: $SRC_BACKEND_FOLDER"
    printInfo " Backend destination folder: $DEST_BACKEND_FOLDER"
    printInfo "     Frontend source folder: $SRC_FRONTEND_FOLDER"
    printInfo "Frontend destination folder: $DEST_FRONTEND_FOLDER"
    printInfo "=================================================="
    printWarn "This script will delete all contents from '$DEST_BACKEND_FOLDER' and"
    printWarn "will replace them with the contents of '$SRC_BACKEND_FOLDER'!"
    printWarn ""
    printWarn "This script will run 'grunt build' on the frontend before"
    printWarn "copying all files into the Dropbox folder!"

    if promptMsg "Are you sure you want to continue?"; then
        buildFrontend
        deploy
    else
        printError "Deployment aborted!"
    fi
}

main
