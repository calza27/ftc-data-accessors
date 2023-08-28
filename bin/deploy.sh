#!/usr/bin/env bash
die() { echo "${1:-ouch}" >&2; exit "${2:-1}"; }

hash sam 2>/dev/null || die "missing dep: sam"

profile=$1
[[ $profile == "" ]] && die "profile needs to be specified"

for path in ./cmd/*/; do
  dirname=$(basename "$path")

  echo "Building $dirname..."

  cd "./cmd/$dirname" || die "failed to cd into ./cmd/$dirname"
  GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o "../bin/$dirname/bootstrap" *.go
  cd ../..
done

sam deploy --no-fail-on-empty-changeset \
--config-file "./samconfig.toml" \
--template-file ./template.yaml \
--profile "$profile"

rm -rf ./cmd/bin # clean up
