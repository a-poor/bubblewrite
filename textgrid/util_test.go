package textgrid

import (
	"reflect"
	"testing"
)

func TestGridify(t *testing.T) {
}

func TestDuplicate(t *testing.T) {
	s := []rune("hello")
	expect := []rune{'h', 'e', 'l', 'l', 'o'}
	got := duplicate(s)
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("Expected %q, got %q", expect, got)
	}
}

func TestDuplicate2D(t *testing.T) {
}

func TestJoin(t *testing.T) {
}

func TestJoin2D(t *testing.T) {
}

func TestSplit(t *testing.T) {
}

func TestSplit2D(t *testing.T) {
}
