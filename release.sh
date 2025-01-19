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
git tag ${VERSION}
git push origin --tags
