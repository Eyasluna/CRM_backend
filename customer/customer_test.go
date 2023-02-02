package customer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = Describe("View Node Test", func() {

	BeforeEach(func() {

	})
	AfterEach(func() {

	})

	Context("when calculating node hash", func() {
		It("should provide different hash if custom attributes are different", func() {

		})

	})
})
