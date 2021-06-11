package logger_test

import (
	"github.com/dzahariev/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("level enumeration", func() {

	Context("check", func() {
		It("all levels contains 7 levels", func() {
			Expect((len(logger.AllLevels()))).To(Equal(7))
		})
	})

})
