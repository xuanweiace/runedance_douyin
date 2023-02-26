#!/usr/bin/env bash
RUN_NAME="user"

mkdir -p output/bin
cp script/* output/
chmod +x output/bootstrap.sh

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    go1.18 build -o output/bin/${RUN_NAME}
else
    go1.18 test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi
