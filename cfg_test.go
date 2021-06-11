package logger_test

import (
	"os"

	"github.com/dzahariev/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("log configuration", func() {
	var config = new(logger.Config)

	BeforeEach(func() {
		config.Default()
	})

	Context("validate", func() {
		It("should be OK when all values are present and valid", func() {
			config.Format = "string"
			config.Level = "info"
			Expect(config.Validate()).ShouldNot(HaveOccurred())
		})

		It("should return error when format is not set", func() {
			config.Format = ""
			Expect(config.Validate()).Should(HaveOccurred())
		})

		It("should return error when level is not set", func() {
			config.Level = ""
			Expect(config.Validate()).Should(HaveOccurred())
		})
	})

	Context("values from environment", func() {
		It("value from environment should be considered with priority", func() {
			envValue := "envValue"
			os.Setenv("LOG_FORMAT", envValue)
			v, err := logger.GetViper()
			Expect(err).ToNot(HaveOccurred())
			config.LoadValues(v)
			Expect(config.Format).To(Equal(envValue))
			os.Unsetenv("LOG_FORMAT")
		})
	})

})
