#!/bin/bash
set -e
set -x

# 引数でバージョンを受け取る
if [ $# -ne 1 ]; then
    echo "バージョンを指定してください"
    exit 1
fi
VERSION=$1
# VERSIONがvで始まっているかチェック
if [[ ! "${VERSION}" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "バージョンはv1.2.3のように指定してください"
    exit 1
fi
ARCH_LIST=("amd64" "arm64")
GOOS="darwin"
for ARCH in "${ARCH_LIST[@]}" ; do
    echo "[ ${ARCH} ]"
    NAME="gcal-run_${GOOS}_${ARCH}_${VERSION}"
    rm -rf "dist/${NAME}"
    mkdir -p "dist/${NAME}"
    GCO_ENABLED=0 GOOS=${GOOS} GOARCH=${ARCH} go build  -ldflags "-X main.version=${VERSION}" -o "dist/${NAME}/gcal_run"  cmd/gcal_run/gcal_run.go
    GCO_ENABLED=0 GOOS=${GOOS} GOARCH=${ARCH} go build  -ldflags "-X main.version=${VERSION}" -o "dist/${NAME}/installer" cmd/installer/installer.go
    (cd dist && tar zcvf ${NAME}.tar.gz ${NAME})
done

