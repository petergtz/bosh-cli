package director_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/cloudfoundry/bosh-cli/director"

	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Director", func() {
	var (
		director Director
		server   *ghttp.Server
	)

	BeforeEach(func() {
		director, server = BuildServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("LatestConfig", func() {
		It("returns latest config if there is at least one", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/configs", "limit=1"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"content": "first"}]`),
				),
			)

			cc, err := director.LatestConfig()
			Expect(err).ToNot(HaveOccurred())
			Expect(cc).To(Equal(Config{Content: "first"}))
		})

		It("returns error if there is no config", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/configs", "limit=1"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[]`),
				),
			)

			_, err := director.LatestConfig()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("No config"))
		})
	})
})
