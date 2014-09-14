package module

import (
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	tree := NewTree(testConfig(t, "basic"))
	actual := tree.Modules()

	expected := []*Module{
		&Module{Name: "foo", Source: "./foo"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("bad: %#v", actual)
	}
}
