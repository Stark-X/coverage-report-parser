package lcov

import (
	"os"
	"strings"
	"testing"
)

func TestParse_emptyFile(t *testing.T) {
	parser := NewParser(strings.NewReader(""))
	ci, err := parser.Parse()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if ci.LineCoverage() != nil {
		t.Errorf("unexpected line coverage: %v", ci.LineCoverage())
	}

	if ci.BranchCoverage() != nil {
		t.Errorf("unexpected branch coverage: %v", ci.BranchCoverage())
	}
	if ci.InstructionCoverage() != nil {
		t.Errorf("unexpected instruction coverage: %v", ci.InstructionCoverage())
	}

	if ci.ClassCoverage() != nil {
		t.Errorf("unexpected class coverage: %v", ci.ClassCoverage())
	}
}

func TestParse_standardFile(t *testing.T) {
	f, err := os.Open("test_data/sample.lcov")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer f.Close()

	parser := NewParser(f)
	ci, err := parser.Parse()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ci.LineCoverage() == nil {
		t.Fatalf("unexpected line coverage: %v", ci.LineCoverage())
	}

	if *ci.LineCoverage() != 0.625 {
		t.Fatalf("unexpected line coverage: %v", ci.LineCoverage())
	}
}
