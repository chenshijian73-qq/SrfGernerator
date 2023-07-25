#!/usr/bin/env bash

set -e

MOD_NAME="github.com/chenshijian73-qq/SrfGernerator"
TARGET_DIR="dist"
TARGET_NAME="srf"
PLATFORMS="darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64"

BUILD_VERSION="2.0.0"
BUILD_DATE=$(date "+%F %T")
COMMIT_SHA1=$(git rev-parse HEAD)

rm -rf ${TARGET_DIR}
mkdir ${TARGET_DIR}

go_install_path=$(echo "${GOPATH}"|awk -F ":" '{print$1}')
install_path="/usr/local/bin"

if [ "$1" == "install" ]; then
  echo "install to ${install_path}"
  go build -o "${go_install_path}"/${TARGET_NAME} -ldflags \
    "-X '${MOD_NAME}/cmd.Version=${BUILD_VERSION}' \
    -X '${MOD_NAME}/cmd.BuildDate=${BUILD_DATE}' \
    -X '${MOD_NAME}/cmd.CommitID=${COMMIT_SHA1}' "
  echo "install => ${install_path}/${TARGET_NAME}"
  ln -sf "${go_install_path}"/${TARGET_NAME} ${install_path}/${TARGET_NAME}
elif [ "$1" == "uninstall" ]; then
    echo "remove => ${go_install_path}/${TARGET_NAME}"
    rm -rf "${go_install_path}"/${TARGET_NAME}
    rm -rf ${install_path}/${TARGET_NAME}
else
  for pl in ${PLATFORMS}; do
      export GOOS=$(echo "${pl}" | cut -d'/' -f1)
      export GOARCH=$(echo "${pl}" | cut -d'/' -f2)
      export CGO_ENABLED=0

      export TARGET=${TARGET_DIR}/${TARGET_NAME}_${GOOS}_${GOARCH}
      if [ "${GOOS}" == "windows" ]; then
          export TARGET=${TARGET_DIR}/${cmd}_${GOOS}_${GOARCH}.exe
      fi

      echo "build => ${TARGET}"
      go build -trimpath -o "${TARGET}" \
              -ldflags    "-X '${MOD_NAME}/cmd.Version=${BUILD_VERSION}' \
                          -X '${MOD_NAME}/cmd.BuildDate=${BUILD_DATE}' \
                          -X '${MOD_NAME}/cmd.CommitID=${COMMIT_SHA1}' \
                          -w -s"
  done
fi