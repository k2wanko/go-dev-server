#!/bin/sh

gox -os "darwin windows" \
    -arch="386 amd64" \
    -output "dist/{{.Dir}}_{{.OS}}_{{.Arch}}"
