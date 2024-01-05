#!/bin/bash
set -e
set -x

# 引数でバージョンを受け取る
if [ $# -ne 1 ]; then
    echo "バージョンを指定してください"
    exit 1
fi
VERSION=$1
ARCH_LIST=("amd64" "arm64")
GOOS="darwin"
for ARCH in "${ARCH_LIST[@]}" ; do
    echo "[ ${ARCH} ]"
    NAME="gcal-run_${GOOS}_${ARCH}_${VERSION}"
    mkdir -p "dist/${NAME}"
    cp README.md "dist/${NAME}/README.md"
    GCO_ENABLED=0 GOOS=${GOOS} GOARCH=${ARCH} go build -o "dist/${NAME}/gcal_run"  cmd/gcal_run/gcal_run.go
    GCO_ENABLED=0 GOOS=${GOOS} GOARCH=${ARCH} go build -o "dist/${NAME}/installer" cmd/installer/installer.go
    (cd dist && tar zcvf ${NAME}.tar.gz ${NAME})
done

