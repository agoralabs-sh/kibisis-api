#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Builds the swagger docs to packages/v1/docs.
#
# Examples
#
#   ./bin/create_swagger_docs.sh
#
# Returns exit code 0.
function main {
  local function_dir
  local go_path
  local package_dir
  local routes

  set_vars

  go_path=$(go env GOPATH)
  routes="./.swag" # the general info must be first

  for package_dir in ./packages/*/; do
    package_dir=${package_dir%*/} # remove the trailing "/"

    for function_dir in "${package_dir}"/*/; do
      function_dir=${function_dir%*/} # remove the trailing "/"

      routes="${routes},${function_dir}"
    done
  done

  "${go_path}"/bin/swag init \
    --dir "${routes}" \
    --output "${PWD}"/packages/v1/docs/src/spec \
    --outputTypes json

  exit 0
}

# And so, it begins...
main
