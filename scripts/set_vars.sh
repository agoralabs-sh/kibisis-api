#!/usr/bin/env bash

# Public: Convenience function that exports some common environment variables.
function set_vars() {
  export BUILD_DIR="${PWD}/.build"
  export CONFIG_DIR="${PWD}/.config"
  export CONFIGS_DIR="${PWD}/configs"
  export CORE_SRC_DIR="${PWD}/cmd/core"
  export ERROR_PREFIX='\033[0;31m[ERROR]\033[0m'
  export INFO_PREFIX='\033[1;33m[INFO]\033[0m'
}
