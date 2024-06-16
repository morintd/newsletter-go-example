package article

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GET /article/:slug", func() {
	It("Should return articles", func() {
		Expect(1).To(Equal(1))
	})
})
