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
	baseURL = os.Getenv("TS_BASE_URL")
	if baseURL == "" {
		baseURL = "tactic.terminus.ninja"
	}
	RunSpecs(t, "E2E Integration Testing Suite")
}
