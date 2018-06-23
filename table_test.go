package turnbased

import (
	"testing"
	"strings"
)

func TestTranscript_Count(t *testing.T) {
	transcript := Transcript("A1B2C3")
	if c := transcript.Count(); c != 3 {
		t.Errorf("expecte 3, got %v", c)
	}
}

func TestTranscript_Count_panic(t *testing.T) {
	transcript := Transcript("A1B2C3D")
	defer func() {
		if recover() == nil {
			t.Fatalf("panic expected")
		}
	}()
	transcript.Count()
}

func TestCellAddress_XY(t *testing.T) {
	for _, ca := range []struct {
		CellAddress
		x, y int
	}{
		{"A1", 0, 0},
		{"B1", 1, 0},
		{"C1", 2, 0},
		{"A2", 0, 1},
		{"B2", 1, 1},
		{"C2", 2, 1},
		{"A3", 0, 2},
		{"B3", 1, 2},
		{"C3", 2, 2},
		{"A12", 0, 11},
		{"B12", 1, 11},
		{"C12", 2, 11},
	}{
		if x := ca.X(); x != ca.x {
			t.Errorf("%v.X():%v != %v", ca.CellAddress, x, ca.x)
		}
		if y := ca.Y(); y != ca.y {
			t.Errorf("%v.Y():%v != %v", ca.CellAddress, y, ca.y)
		}
	}
}

func TestNewSize(t *testing.T) {
	t.Run("zero_width", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("panic expected")
			} else if !strings.Contains(r.(string), "width") {
				t.Errorf("panic regards width expected")
			}
		}()
		NewSize(0, 1)
	})
	t.Run("zero_height", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("panic expected")
			} else if !strings.Contains(r.(string), "height") {
				t.Errorf("panic regards height expected")
			}
		}()
		NewSize(1, 0)
	})
	t.Run("ok_sizes", func(t *testing.T) {
		testSize := func(w, h int, s string) {
			t.Helper()
			if size := NewSize(w, h); string(size) != s {
				t.Errorf("%d:%d expects %v, got: %v", w, h, s, size)
			}
		}
		testSize(1,1,"A1")
		testSize(1,2,"A2")
		testSize(1,3,"A3")
		testSize(2,1,"B1")
		testSize(2,2,"B2")
		testSize(2,3,"B3")
		testSize(3,1,"C1")
		testSize(3,2,"C2")
		testSize(3,3,"C3")
	})
}