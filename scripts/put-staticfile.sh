#!/usr/bin/env bash

set -euo pipefail

if [[ $# != 1 ]]; then
    echo "Usage: ./scripts/put-staticfile.sh /path/to/target-file"
    exit 1
fi

TARGET_PATH=$1
BASE_DIR="./statics/public"

TARGET_FILENAME=$(basename $TARGET_PATH)
TARGET_HASH=$(openssl dgst -sha256 $TARGET_PATH | cut -d' ' -f2)
NEW_TARGET_PATH="${BASE_DIR}/${TARGET_HASH}"

mkdir -p $NEW_TARGET_PATH
mv -v $TARGET_PATH $NEW_TARGET_PATH/


