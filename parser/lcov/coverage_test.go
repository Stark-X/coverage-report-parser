package lcov

import "testing"

func TestAddFrom(t *testing.T) {
	ci := &CoverageItem{
		LineFound:   1,
		LineHint:    2,
		BranchFound: 3,
		BranchHint:  4,
	}
	other := &CoverageItem{
		LineFound:   5,
		LineHint:    6,
		BranchFound: 7,
		BranchHint:  8,
	}

	ci.AddFrom(other)

	if ci.LineFound != 6 {
		t.Errorf("LineFound: got %d, want %d", ci.LineFound, 6)
	}
	if ci.LineHint != 8 {
		t.Errorf("LineHint: got %d, want %d", ci.LineHint, 8)
	}
	if ci.BranchFound != 10 {
		t.Errorf("BranchFound: got %d, want %d", ci.BranchFound, 10)
	}
	if ci.BranchHint != 12 {
		t.Errorf("BranchHint: got %d, want %d", ci.BranchHint, 12)
	}
}
