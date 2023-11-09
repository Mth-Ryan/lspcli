RUNTIME="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
BIN=${RUNTIME}/bin

export PATH="${PATH}:${BIN}"
