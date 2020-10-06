package hashtable

import (
	"github.com/wednesday-solutions/golang-algorithm-club/pkg/utl"
	"testing"
)

func TestMain(m *testing.M) {
	utl.SetupTests(m, insertKeys)
}

func insertKeys() {
	Insert("GOLANG", "001")
	Insert("LANGGO", "100")
	Insert("GO", "002")
	Insert("HASH", "003")
}

func TestGetValueFromBucket(t *testing.T) {
	// tests for GetValueFromBucket
	golangValue := GetValueFromBucket("GOLANG")
	if golangValue != "001" {
		t.Errorf("Value for GOLANG should be 001, got value = %v", golangValue)
	}
	langgoValue := GetValueFromBucket("LANGGO")
	if langgoValue != "100" {
		t.Errorf("Value for GOLANG should be 100, got value = %v", langgoValue)
	}
	goValue := GetValueFromBucket("GO")
	if goValue != "002" {
		t.Errorf("Value for GOLANG should be 002, got value = %v", goValue)
	}
	hashValue := GetValueFromBucket("HASH")
	if hashValue != "003" {
		t.Errorf("Value for GOLANG should be 003, got value = %v", hashValue)
	}
}
