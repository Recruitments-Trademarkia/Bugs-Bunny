#!/usr/bin/env bash

# Check gofmt
echo "==> Checking that code complies with gofmt requirements..."
gofmt_files=$(find . -name '*.go' | grep -v vendor | xargs gofmt -l -s)

IFS=$'\n' read -ra ADDR -d $'\0' <<< "$gofmt_files"

for i in "${ADDR[@]}"
do
    echo "gofmt working on following file:  $i"
    echo $(gofmt -w $i)
done

exit 0
