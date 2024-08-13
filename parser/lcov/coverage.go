package lcov

type CoverageItem struct {
	SourcePath  string
	LineFound   int
	LineHint    int
	BranchFound int
	BranchHint  int
}

func (ci *CoverageItem) AddFrom(other *CoverageItem) {
	ci.LineFound += other.LineFound
	ci.LineHint += other.LineHint
	ci.BranchFound += other.BranchFound
	ci.BranchHint += other.BranchHint
}
