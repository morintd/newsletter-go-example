#!/bin/bash

# Find all __tests__ directories
TEST_DIRS=$(find . -type d -name '__tests__')

ginkgo $TEST_DIRS