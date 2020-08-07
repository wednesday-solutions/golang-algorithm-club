package utl

import (
	"os"
	"testing"
)

//SetupTests Setup the tests
func SetupTests(m *testing.M, fn func()) {
	fn()
	code := m.Run()
	os.Exit(code)
}
