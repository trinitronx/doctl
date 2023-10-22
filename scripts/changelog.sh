#!/usr/bin/env bash

set -o pipefail

tfile=$(mktemp /tmp/doctl-CHANGELOG-XXXXXX)
github-changelog-generator -org trinitronx -repo doctl >"$tfile"

GO111MODULE=on go mod tidy

echo "$tfile"
