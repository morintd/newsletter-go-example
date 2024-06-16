#!/bin/bash

TEST_DIRS=$(find . -type d -name '__tests__')

# @TODO: won't reload directories created after starting command ? :(
ginkgo watch $TEST_DIRS