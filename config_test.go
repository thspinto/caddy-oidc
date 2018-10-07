package oidc

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCaddyOidcConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CaddyOidc Config Suite")
}

var _ = Describe("OIDC Config", func() {
	Describe("OIDC config block", func() {
		It("parses configuration", func() {

		})
	})
})
