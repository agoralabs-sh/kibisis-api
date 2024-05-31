#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Iterates through each function and runs the tests.
#
# Examples
#
#   ./bin/test.sh
#
# Returns exit code 0.
function main {
  local function_dir
  local package_dir

  set_vars

  for package_dir in ./packages/*/; do
    for function_dir in "${package_dir}"/*/; do
      # remove the trailing "/"
      function_dir=${function_dir%*/}

      # change directory to the function directory and run the tests
      (cd "${function_dir}" && go test -v)
    done
  done

  exit 0
}

# And so, it begins...
main
