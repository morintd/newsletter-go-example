# newsletter-go-example

Implementation of a newsletter API in Go using [Fx](https://github.com/uber-go/fx) and [Gin](https://github.com/gin-gonic/gin).

Tests with [Ginkgo](https://github.com/onsi/ginkgo), using an uncommon pattern with `__tests__` (Influenced by `Node`/`Jest`).

No database, but an in-memory implementation of repositories.

## Commands

```sh
$ make dev
```

```sh
$ make test
```

```sh
$ make test_watch
```