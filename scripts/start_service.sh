#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Starts a service from the services/ directory.
#
# $1 - the name of the service.
#
# Examples
#
#   ./bin/start_service.sh "relay"
#
# Returns exit code 0.
function main() {
  set_vars

  printf "%b starting ${1} service...\n" "${INFO_PREFIX}"
  (cd services/${1} && \
    CompileDaemon \
    -build="go build -o ${BUILD_DIR}/${1} cmd/${1}/main.go" \
    -color=true \
    -command="${BUILD_DIR}/${1}")

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main "$@"
