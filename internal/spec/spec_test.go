package spec_test

import (
	"path/filepath"
	"runtime"

	. "github.com/genesis-community/testkit/testing"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Interal Kit", func() {
	BeforeSuite(func() {
		_, filename, _, _ := runtime.Caller(0)
		KitDir, _ = filepath.Abs(filepath.Join(filepath.Dir(filename), "../"))
	})

	Test(Environment{
		Name:        "baseline",
		CloudConfig: "aws",
	})
	Test(Environment{
		Name:        "openvpn",
		CloudConfig: "aws",
	})
	Test(Environment{
		Name:        "provided-user",
		CloudConfig: "aws",
	})
	Test(Environment{
		Name:        "detect-cpi",
		CloudConfig: "aws",
		CPI:         "aws",
	})
	Test(Environment{
		Name:        "bosh-variables",
		CloudConfig: "aws",
	})
	Test(Environment{
		Name:        "credhub",
		CloudConfig: "aws",
	})
})
