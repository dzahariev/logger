package logger_test

import (
	"github.com/dzahariev/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("formatter enumeration", func() {

	Context("check", func() {
		It("all formatters contains 2 variations", func() {
			Expect((len(logger.AllFormatters()))).To(Equal(2))
		})
	})

})
