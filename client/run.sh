#!/bin/bash
scriptdir="$(dirname "$(readlink -f "$0")")"
echo "scriptdir: $scriptdir"
cd $scriptdir

go run .
