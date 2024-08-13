package lcov

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"coverage-report-parser/parser"
)

type Parser struct {
	reader io.Reader
}

func NewParser(r io.Reader) *Parser {
	return &Parser{reader: r}
}

func (p *Parser) Parse() (*parser.CoverageInfo, error) {
	lineTotal := 0
	lineCovered := 0
	branchCovered := 0
	branchTotal := 0

	items, err := p.ParseLcov()
	if err != nil {
		err = fmt.Errorf("parsing failed %w", err)
		return nil, err
	}

	for _, item := range items {
		lineTotal += item.LineFound
		lineCovered += item.LineHint
		branchTotal += item.BranchFound
		branchCovered += item.BranchHint
	}

	return parser.NewCoverageInfoBuilder().
		WithLineCovered(lineCovered).
		WithLineTotal(lineTotal).
		WithBranchCovered(branchCovered).
		WithBranchTotal(branchTotal).
		Build(), nil
}

func (p *Parser) ParseLcov() ([]*CoverageItem, error) {
	memory := make(map[string]*CoverageItem)

	scanner := bufio.NewScanner(p.reader)
	coverageItem := &CoverageItem{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		err := error(nil)
		switch {
		case strings.HasPrefix(line, "SF:"):
			coverageItem.SourcePath = strings.TrimPrefix(line, "SF:")
		case strings.HasPrefix(line, "LH:"):
			coverageItem.LineHint, err = strconv.Atoi(strings.TrimPrefix(line, "LH:"))
		case strings.HasPrefix(line, "LF:"):
			coverageItem.LineFound, err = strconv.Atoi(strings.TrimPrefix(line, "LF:"))
		case strings.HasPrefix(line, "BRH:"):
			coverageItem.BranchHint, err = strconv.Atoi(strings.TrimPrefix(line, "BRH:"))
		case strings.HasPrefix(line, "BRF:"):
			coverageItem.BranchFound, err = strconv.Atoi(strings.TrimPrefix(line, "BRF:"))
		case strings.HasPrefix(line, "end_of_record"):
			if exists, ok := memory[coverageItem.SourcePath]; ok {
				exists.AddFrom(coverageItem)
			} else {
				memory[coverageItem.SourcePath] = coverageItem
				coverageItem = &CoverageItem{}
			}
		}
		if err != nil {
			err = fmt.Errorf("parsing failed %w", err)
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		err = fmt.Errorf("reading content failed %w", err)
		return nil, err
	}

	result := make([]*CoverageItem, 0, len(memory))
	for _, item := range memory {
		result = append(result, item)
	}

	return result, nil
}
