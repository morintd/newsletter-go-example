TEST_DIRS := $(shell find . -type d -name '__tests__')

dev:
	gin
test:
	"$(CURDIR)/scripts/test.sh"
test_watch:
	"$(CURDIR)/scripts/test_watch.sh"

.NOTPARALLEL:

.PHONY: dev test test_watch