# coverage-report-parser

Parser for coverage report. extract total coverage info from the coverage report. writing in golang.

## Usage

`go run main.go --report ./parser/lcov/test_data/sample.lcov`

or use it after build it `make build && ./dist/coverage-report-parse --report ./parser/lcov/test_data/sample.lcov`

## RoadMap

- [x] support lcov file parsing
- [ ] support jacococ xml report parsing
