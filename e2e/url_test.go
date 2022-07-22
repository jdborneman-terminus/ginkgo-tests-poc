package e2e

import (
	"fmt"
	"log"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Try Google for a 200", func() {
	It("should return 200", func() {
		log.Println("------------------------>SLEEP NOW<------------------------")
		time.Sleep(4 * time.Second)
		resp, requestErr := http.Get(fmt.Sprintf("https://%s", baseURL))
		Expect(requestErr).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))
	})
})
