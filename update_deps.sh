#!/bin/bash

bzl run @go_sdk//:bin/go -- get -v ./...
bzl run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
bzl run //:gazelle
