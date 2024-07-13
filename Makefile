
dev:
	gin
test:
	"$(CURDIR)/scripts/test.sh"
test_watch:
	"$(CURDIR)/scripts/test_watch.sh"

.NOTPARALLEL:

.PHONY: dev test test_watch