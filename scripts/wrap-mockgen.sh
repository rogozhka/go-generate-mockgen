#!/usr/bin/env bash
#
#//go:generate wrap-mockgen.sh -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
#
# This script is intended to be called from go:generate tightly with $GOLINE, $GOFILE, $GOPACKAGE variables.
# It must be available in $PATH and have exec permission (chmod +x).
# After preparing arguments, finally calls go.uber.org/mock/mockgen.
# Installation: go install go.uber.org/mock/mockgen@latest
#
# For usage:
# place the go:generate call on some line before(!) the declaration of mocked interface.
# The closest line found with the interface declaration will be used for mock generation,
# follow the arguments format: wrap-mockgen.sh -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
#
# EXAMPLE:
# //go:generate wrap-mockgen.sh -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
# type treeStoredItemInterface interface {
#
# privateInterfaceName will be transformed into MockPrivateInterfaceName w/ all the methods mocked.
#
# By default all the mocks are stored in the nested "mocks" package.
# You can edit this by overriding TARGET_PACKAGE variable:
# for example, to generate in the current package change to
# TARGET_PACKAGE=$PACKAGE
#
# if you prefer packagename_mocks style, use
# TARGET_PACKAGE=${PACKAGE}_mocks
#
# Happy mocking! :)
#
# PS/ you are free to modify this script, released under MIT Licence
# Check https://github.com/rogozhka/go-generate-mockgen for usage example w/ couple of tests.
# Copyright (c) 2024 Serge R. thecoder@yandex.ru

LINE=""
SOURCE=""
PACKAGE=""

for arg in "$@"
do
    case $arg in
        -line=*)
            LINE="${arg#-line=}"
            ;;
        -source=*)
            SOURCE="${arg#-source=}"
            ;;
        -package=*)
            PACKAGE="${arg#-package=}"
            ;;
        *)
            echo "Unknown argument: $arg"
            exit 1
            ;;
    esac
done

TARGET_LINE_CONTENT=$(tail -n +$LINE "$SOURCE" | grep -m 1 '^type ')

INTERFACE=$(echo "$TARGET_LINE_CONTENT" | awk '{ for(i=1; i<=NF; i++) if ($i == "type") print $(i+1) }')

#capitalize
MOCK_NAME=$(echo "${INTERFACE:0:1}" | tr '[:lower:]' '[:upper:]')"${INTERFACE:1}"
MOCK_NAME="Mock"${MOCK_NAME}

TARGET=$(basename "$SOURCE" .go)_generated.go

TARGET_PACKAGE=mocks
# optional: for mypackage_mocks
#TARGET_PACKAGE=${PACKAGE}_mocks

GOMOD_CHECK_PATH="./go.mod"

if [ ! -f $GOMOD_CHECK_PATH  ]; then
    # In case of run from docker.
    # The go.mod file is missing from the directory. Since go:generate runs the script
    # in the current directory of the file, we can only Docker-mount the directory
    # containing the file, while go.mod might be several levels above.
    # Therefore, let's create a go.mod-stub to prevent mockgen
    # from going crazy searching for it.
    echo "module $PACKAGE" > $GOMOD_CHECK_PATH
    echo "go 1.23.4" >> $GOMOD_CHECK_PATH
    go_mod_stub_was_created=true
fi

mockgen -source=$SOURCE \
-destination=${TARGET_PACKAGE}/$TARGET \
-mock_names ${INTERFACE}=${MOCK_NAME} \
-typed \
-package=$TARGET_PACKAGE \
$INTERFACE

if [ "$go_mod_stub_was_created" = true ]; then
    rm $GOMOD_CHECK_PATH
fi
