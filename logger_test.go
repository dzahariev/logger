package logger_test

import (
	"testing"

	"github.com/dzahariev/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestForLogger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logger Suite")
}

var _ = Describe("loger tests", func() {
	Context("init", func() {
		It("should be OK when all values are present and valid", func() {
			config := new(logger.Config)
			config.Format = "string"
			config.Level = "info"
			err := logger.InitLogger(*config)
			Expect(err).ToNot(HaveOccurred())
		})
	})

})
