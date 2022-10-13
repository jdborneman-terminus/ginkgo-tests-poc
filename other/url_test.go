package other

import (
	"fmt"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Try URL for a 200", func() {
	It("should return 200", func() {
		time.Sleep(60 * time.Second)
		resp, requestErr := http.Get(fmt.Sprintf("https://%s", baseURL))
		Expect(requestErr).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))
	})
})
