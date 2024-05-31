#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Creates a .env file if it don't exist.
#
# Examples
#
#   ./scripts/setup.sh
#
# Returns exit code 1 if no example file exists, otherwise, exit code 0 is returned.
function main() {
  local env_example_file_path
  local env_file_path

  set_vars

  env_example_file_path="${PWD}/.env.example"

  if [[ ! -f "${env_example_file_path}" ]];
    then
      printf "%b no .env example at %b \n" "${ERROR_PREFIX}" "${env_example_file_path}"
      exit 1
  fi

  env_file_path="${PWD}/.env"

  printf "%b creating %b files...\n" "${INFO_PREFIX}" "${env_file_path}"

  # create the .env file
  cp -n "${env_example_file_path}" "${env_file_path}"

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main
