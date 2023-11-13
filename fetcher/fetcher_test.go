package fetcher_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"os"

	"github.com/onsi/gomega/ghttp"
	"gitlab.com/oozie/example/fetcher"
)

const sampleResponse = `{
  "name" : "test-name",
  "cluster_name" : "test-cluster",
  "cluster_uuid" : "test-uuid",
  "version" : {
    "number" : "8.4.3",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "42f05b9372a9a4a470db3b52817899b99a76ee73",
    "build_date" : "2022-10-04T07:17:24.662462378Z",
    "build_snapshot" : false,
    "lucene_version" : "9.3.0",
    "minimum_wire_compatibility_version" : "7.17.0",
    "minimum_index_compatibility_version" : "7.0.0"
  },
  "tagline" : "You Know, for Search"
}`

var _ = Describe("Fetcher", func() {
	var (
		server *ghttp.Server
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		os.Setenv("ELASTIC_ADDR", server.Addr())

		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/"),
				ghttp.RespondWith(200, sampleResponse),
			),
		)
	})

	It("fetches and parses elasticsearch greeting", func() {
		response, err := fetcher.GetGreetingFromURL("http://" + os.Getenv("ELASTIC_ADDR"))
		Expect(err).NotTo(HaveOccurred())

		Expect(response.ClusterName).To(Equal("test-cluster"))
		Expect(response.ClusterUuid).To(Equal("test-uuid"))
		Expect(response.Tagline).To(Equal("You Know, for Search"))
		Expect(string(response.VersionInfo.LuceneVersion)).To(Equal("9.3.0"))
		Expect(response.VersionInfo.BuildSnapshot).To(BeFalse())
	})

})
