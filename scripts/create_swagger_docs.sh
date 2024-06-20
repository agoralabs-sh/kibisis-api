#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}"/set_vars.sh

# Public: Builds the swagger.json to be served from packages/vX/swagger. Use the argument to direct to which major
# version of the API to create the swagger JSON document for.
#
# Examples
#
#   ./bin/create_swagger_docs.sh 1
#
# Returns exit code 0.
function main {
  local function_dir
  local go_path
  local package_dir
  local routes

  set_vars

  if [ -z "${1}" ]; then
    printf "%b no major version specified, use: ./bin/create_swagger_docs.sh [major_version] \n" "${ERROR_PREFIX}"
    exit 1
  fi

  go_path=$(go env GOPATH)
  package_dir="./packages/v${1}"
  routes="${package_dir}/swagger" # the general info file must be first

  for function_dir in "${package_dir}"/*/; do
    function_dir=${function_dir%*/} # remove the trailing "/"

    # ignore the directory where general info file is, we have already added it
    if [[ "${function_dir}" == *"/v${1}/swagger"* ]]; then
      continue
    fi

    routes="${routes},${function_dir}"
  done

  echo $routes

  "${go_path}"/bin/swag init \
    --dir "${routes}" \
    --output "${PWD}/packages/v${1}/swagger" \
    --outputTypes json

  exit 0
}

# and so, it begins...
main "$@"
