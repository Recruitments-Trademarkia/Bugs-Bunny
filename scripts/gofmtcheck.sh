#!/usr/bin/env bash

# Check gofmt
echo "==> Checking that code complies with gofmt requirements..."
gofmt_files=$(find . -name '*.go' | grep -v vendor | xargs gofmt -l -s)
if [[ -n ${gofmt_files} ]]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    exit 1
fi
echo "All fine! "

exit 0
