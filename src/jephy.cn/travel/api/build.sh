#!/bin/sh -e

sudo dd if=/dev/zero of=/swapfile bs=64M count=16
sudo mkswap /swapfile
sudo swapon /swapfile

echo "[BUILD_INFO] Build start from dir: $(pwd)"
export GO_VERSION="1.10.3"
export PROJECT_NAME="travel_api"
export CMD_FOLDER="cmd"
export BUILD_ROOT=$(pwd)/../
##################  Make a standard go build environment  #####################
export GOPATH=$(pwd)/../../../../
export GOBIN=$BUILD_ROOT/bin

mkdir -p $BUILD_ROOT/pkg $BUILD_ROOT/bin

##########################  go install  #######################################
export PATH=/usr/local/go$GO_VERSION/go/bin:$PATH
alias GO=go
export BUILD_RES=0
echo "[BUILD_INFO] Start build with go insatll."
if [ $# = 0 ] ; then
    CGO_ENABLED=0 GO install ./$CMD_FOLDER/$PROJECT_NAME
    BUILD_RES=$?

    if [[ $BUILD_RES = 0 ]]; then
       echo "[BUILD_INFO] Build success......"
    fi
    exit $BUILD_RES
fi

for i in "$@" ;
do
    CGO_ENABLED=0 GO install ./$CMD_FOLDER/$i
    BUILD_RES=$?
    if [[ $BUILD_RES != 0 ]]; then
        break
    fi
done

if [[ $BUILD_RES = 0 ]]; then
    echo "[BUILD_INFO] Build success......"
fi
exit $BUILD_RES

sudo swapoff /swapfile
sudo rm /swapfile
