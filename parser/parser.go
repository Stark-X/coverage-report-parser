package parser

import (
	"strconv"
)

type CoverageInfo struct {
	InstructionCovered *int
	InstructionTotal   *int
	BranchCovered      *int
	BranchTotal        *int
	LineCovered        *int
	LineTotal          *int
	ClassCovered       *int
	ClassTotal         *int
}

func (ci *CoverageInfo) String() string {
	result := "CoverageInfo{"
	appendField := func(name string, value *int) {
		if value != nil {
			result += name + ": " + strconv.Itoa(*value) + ", "
		}
	}

	appendField("InstructionCovered", ci.InstructionCovered)
	appendField("InstructionTotal", ci.InstructionTotal)
	appendField("BranchCovered", ci.BranchCovered)
	appendField("BranchTotal", ci.BranchTotal)
	appendField("LineCovered", ci.LineCovered)
	appendField("LineTotal", ci.LineTotal)
	appendField("ClassCovered", ci.ClassCovered)
	appendField("ClassTotal", ci.ClassTotal)

	appendCoverage := func(name string, coverageFunc func() *float64) {
		if coverage := coverageFunc(); coverage != nil {
			result += name + ": " + strconv.FormatFloat(*coverage, 'f', 4, 64) + ", "
		}
	}

	appendCoverage("LineCoverage", ci.LineCoverage)
	appendCoverage("BranchCoverage", ci.BranchCoverage)
	appendCoverage("InstructionCoverage", ci.InstructionCoverage)
	appendCoverage("ClassCoverage", ci.ClassCoverage)

	return result + "}"
}

func (ci *CoverageInfo) coveraged(covered, total *int) *float64 {
	if total == nil || covered == nil || *total == 0 {
		return nil
	}
	res := float64(*covered) / float64(*total)
	return &res
}

func (ci *CoverageInfo) LineCoverage() *float64 {
	return ci.coveraged(ci.LineCovered, ci.LineTotal)
}

func (ci *CoverageInfo) BranchCoverage() *float64 {
	return ci.coveraged(ci.BranchCovered, ci.BranchTotal)
}

func (ci *CoverageInfo) InstructionCoverage() *float64 {
	return ci.coveraged(ci.InstructionCovered, ci.InstructionTotal)
}

func (ci *CoverageInfo) ClassCoverage() *float64 {
	return ci.coveraged(ci.ClassCovered, ci.ClassTotal)
}

type CoverageInfoBuilder struct {
	info *CoverageInfo
}

func NewCoverageInfoBuilder() *CoverageInfoBuilder {
	return &CoverageInfoBuilder{info: &CoverageInfo{}}
}

func (b *CoverageInfoBuilder) Build() *CoverageInfo {
	return b.info
}

func (b *CoverageInfoBuilder) WithInstructionCovered(instructionCovered int) *CoverageInfoBuilder {
	b.info.InstructionCovered = &instructionCovered
	return b
}

func (b *CoverageInfoBuilder) WithInstructionTotal(instructionTotal int) *CoverageInfoBuilder {
	b.info.InstructionTotal = &instructionTotal
	return b
}

func (b *CoverageInfoBuilder) WithBranchCovered(branchCovered int) *CoverageInfoBuilder {
	b.info.BranchCovered = &branchCovered
	return b
}

func (b *CoverageInfoBuilder) WithBranchTotal(branchTotal int) *CoverageInfoBuilder {
	b.info.BranchTotal = &branchTotal
	return b
}

func (b *CoverageInfoBuilder) WithLineCovered(lineCovered int) *CoverageInfoBuilder {
	b.info.LineCovered = &lineCovered
	return b
}

func (b *CoverageInfoBuilder) WithLineTotal(lineTotal int) *CoverageInfoBuilder {
	b.info.LineTotal = &lineTotal
	return b
}

func (b *CoverageInfoBuilder) WithClassCovered(classCovered int) *CoverageInfoBuilder {
	b.info.ClassCovered = &classCovered
	return b
}

func (b *CoverageInfoBuilder) WithClassTotal(classTotal int) *CoverageInfoBuilder {
	b.info.ClassTotal = &classTotal
	return b
}

type Parser interface {
	Parse() (*CoverageInfo, error)
}
