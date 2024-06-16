package main

import (
	"newsletter/app/core"

	_ "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/gomega"
)

func main() {
	r := core.Bootstrap()
	r.Run()
}
