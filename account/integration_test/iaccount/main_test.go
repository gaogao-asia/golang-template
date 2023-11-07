package iaccount

import (
	"os"
	"testing"

	itest "github.com/gaogao-asia/golang-template/integration_test"
)

func TestMain(m *testing.M) {
	itest.Setup()

	ret := m.Run()

	os.Exit(ret)
}
