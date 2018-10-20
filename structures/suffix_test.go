package structures

import (
	"testing"
	"reflect"
)

func TestSuffixArray(t *testing.T) {

	testStruct := &SuffixArray{}
	testStruct.SetString("daavidatsadwfgawdadavvid")

	expected := 24
	got := testStruct.GetSuffixArray()

	if len(got) != expected {
		t.Errorf("wrong result #1")
	}

	gotSecond := testStruct.Search("da")
	expectedSecond := []string{
		"daavidatsadwfgawdadavvid",
		"dadavvid",
		"datsadwfgawdadavvid",
		"davvid",
	}
	if !reflect.DeepEqual(expectedSecond, gotSecond) {
		t.Errorf("wrong result #2")
	}
}
