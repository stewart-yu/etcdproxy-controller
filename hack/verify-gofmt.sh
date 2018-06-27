#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# gofmt
bad_files=$(echo $PKGS | xargs $GOFMT -l)
if [[ -n "${bad_files}" ]]; then
  echo "✖ gofmt needs to be run on the following files: "
  echo "${bad_files}"
  exit 1
fi