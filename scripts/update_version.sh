#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Updates the ./packages/system/versions/VERSION file with the supplied version.
#
# $1 - the version to update
#
# Examples

#   ./scripts/update_version_file.sh "1.2.3"
#
# Returns exit code 0 if successful, or 1 if the semantic version is incorrectly formatted.
function main() {
  set_vars

  if [ -z "${1}" ]; then
    printf "%b no version specified, use: ./bin/update_version_file.sh [version] \n" "${ERROR_PREFIX}"
    exit 1
  fi

  # check the input is in semantic version format
  if [[ ! "${1}" =~ ^[0-9]+\.[0-9]+\.[0-9]+ ]]; then
    printf "%b invalid semantic version, got '${1}', but should be in the format '1.0.0' \n" "${ERROR_PREFIX}"
    exit 1
  fi

  # remove the previous contents
  true > "${PWD}"/packages/system/versions/VERSION

  # use the new version
  echo "$1" >> "${PWD}"/packages/system/versions/VERSION

  printf "%b new version set to: %b\n" "${INFO_PREFIX}" "$1"

  exit 0
}

# and so, it begins...
main "$@"
