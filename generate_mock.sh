#!/bin/bash

# generate_mock is responsible for creating mocks from interfaces throughout
# the framework. It must receive the following arguments:
#
# - the internal package path;
# - the file name which will be handled inside this package;
# - the output module name.
generate_mock() {
    local path=$1
    local filename=$2
    local module=$3

    mockgen -source=$path/$filename.go -destination=$path/mock/$filename.go -package mock_$module
    echo "Mocks for $module generated"
}

# Generate mocks from container interfaces
generate_mock "internal/example" "example" "example"