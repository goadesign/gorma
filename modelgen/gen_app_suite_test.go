package modelgen_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestModelGen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ModelGen Suite")
}
