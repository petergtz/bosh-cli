package director_test

import (
	"net/http"

	. "github.com/cloudfoundry/bosh-cli/director"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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
					ghttp.VerifyRequest("GET", "/configs/my-type", "name=&limit=1"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"content": "first"}]`),
				),
			)

			cc, err := director.LatestConfig("my-type", "")
			Expect(err).ToNot(HaveOccurred())
			Expect(cc).To(Equal(Config{Content: "first"}))
		})

		It("returns error if there is no config", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/configs/my-type", "name=&limit=1"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[]`),
				),
			)

			_, err := director.LatestConfig("my-type", "")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("No config"))
		})

		It("returns error if there is no config", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/configs/fake-type", "name=&limit=1"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[]`),
				),
			)

			_, err := director.LatestConfig("fake-type", "")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("No config"))
		})

		It("returns error if info response in non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/configs/fake-type"), server)

			_, err := director.LatestConfig("fake-type", "")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding configs: Director responded with non-successful status code"))
		})
	})

	Describe("UploadConfig", func() {
		It("updates config", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/configs/fake-type", "name="),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type": []string{"text/yaml"},
					}),
					ghttp.RespondWith(http.StatusOK, `{}`),
				),
			)

			err := director.UpdateConfig("fake-type", "", []byte("---"))
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if response is not 200", func() {
			AppendBadRequest(ghttp.VerifyRequest("POST", "/configs/fake-type"), server)

			err := director.UpdateConfig("fake-type", "", nil)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Updating config: Director responded with non-successful status code"))
		})
	})
})
