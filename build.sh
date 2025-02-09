#!/bin/bash

# 引数でバージョンを受け取る
if [ $# -ne 1 ]; then
    echo "バージョンを第一引数に指定してください"
    echo "usage: bash build.sh v1.2.3"
    exit 1
fi

set -x
set -e

VERSION=$1
# VERSIONがvで始まっているかチェック
if [[ ! "${VERSION}" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "バージョンはv1.2.3のように指定してください"
    exit 1
fi

APP_LIST=("gcal_run" "installer")

dataArray=(
    'darwin amd64'
    'darwin arm64'
    'windows amd64'
)

# スクリプトのあるディレクトリを変数に格納
SCRIPT_DIR=$(cd $(dirname $0); pwd)

for i in "${dataArray[@]}"; do
    data=(${i[@]})
    GOOS=${data[0]}
    ARCH=${data[1]}

    echo "[ ${ARCH} ]"
    NAME="gcal-run_${GOOS}_${ARCH}_${VERSION}"
    rm -rf "dist/${NAME}"
    mkdir -p "dist/${NAME}"

    for APP in "${APP_LIST[@]}"; do
        cd "${SCRIPT_DIR}/cmd/${APP}/"
        if [ "${GOOS}" = "windows" ]; then
            # Windowsの場合はアイコンファイルを埋め込む
            cp ${SCRIPT_DIR}/resource/gcal_run.syso ${SCRIPT_DIR}/cmd/gcal_run/
            APP="${APP}.exe"
        fi

        GCO_ENABLED=0 GOOS=${GOOS} GOARCH=${ARCH} go build \
            -ldflags "-X main.version=${VERSION}" \
            -o "${SCRIPT_DIR}/dist/${NAME}/${APP}"

        if [ "${GOOS}" = "windows" ]; then
            rm ${SCRIPT_DIR}/cmd/gcal_run/gcal_run.syso
        fi
    done

    cd "${SCRIPT_DIR}/dist"
    tar zcvf ${NAME}.tar.gz ${NAME}
done

cd "${SCRIPT_DIR}"