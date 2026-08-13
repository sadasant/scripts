#!/bin/sh
set -eu

cd "$(dirname "$0")"
OPENAI_HTTP_TRANSPORT="${OPENAI_HTTP_TRANSPORT:-curl}"
export OPENAI_HTTP_TRANSPORT
exec go run . "$@"
