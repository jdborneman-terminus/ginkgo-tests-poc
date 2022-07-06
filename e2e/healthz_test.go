package e2e

import (
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Healthz Route", func() {
	It("should return 200", func() {
		resp, requestErr := http.Get(fmt.Sprintf("https://%s/v1/ttd/all/healthz", baseURL))
		Expect(requestErr).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))
	})
})
