package e2e

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var baseURL string

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	baseURL = os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "google.com"
	}
	RunSpecs(t, "E2E Integration Testing Suite")
}
