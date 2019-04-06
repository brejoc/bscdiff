// bscdiff compares bsc, issue and CVE numbers from a source changelog, to a
// target changelog. Missing numbers are then printed with their occurrence
// in the source changelog.
//
// Usage: bscdiff <source_file> <target_file>

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

type searchResult struct {
	line  int
	match []string
	text  string
}

var regexStrings []string

func main() {
	// These are the regex-strings that will be used later on.
	regexStrings = []string{
		`bsc#\d*`,
		`U#\d*`,
		`(CVE-(1999|2\d{3})-(0\d{2}[1-9]|[1-9]\d{3,}))`}

	args := os.Args

	if len(args) > 1 && (args[1] == "-h" || args[1] == "--help") {
		fmt.Println("bscdiff compares bsc, issue and CVE numbers from a source changelog, ")
		fmt.Println("to a target changelog. Missing numbers are then printed with their ")
		fmt.Println("occurrence in the source changelog")
		fmt.Println()
		fmt.Println(fmt.Sprintf("usage: %s <source_file> <target_file>\n", args[0]))
		os.Exit(0)
	}

	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: %s <source_file> <target_file>", args[0])
		os.Exit(1)
	}

	// Check if files are actually there… and files.
	for _, file := range os.Args[1:3] {
		if !fileExists(file) {
			fmt.Fprintf(os.Stderr, "%s does not exist!", file)
			os.Exit(1)
		}
	}

	searchResults1 := scanFile(args[1])
	searchResults2 := scanFile(args[2])
	missingBscs := findMissingBsc(searchResults1, searchResults2)
	prettyPrintMissingBscs(searchResults1, missingBscs)
}

// Outputs the missing BSC numbers in a useful format.
func prettyPrintMissingBscs(searchResults1 []searchResult, missingBscs []string) {
	sort.Strings(missingBscs)
	for _, bsc := range missingBscs {
		for _, searchResult := range searchResults1 {
			searchPos := sort.SearchStrings(searchResult.match, bsc)
			if searchPos < len(searchResult.match) && searchResult.match[searchPos] == bsc {
				fmt.Println(fmt.Sprintf("%d: %s -> %s",
					searchResult.line,
					bsc,
					searchResult.text))
			}
		}
	}
}

// Returns a list of BSC numbers, that are missing from the second changelog file.
func findMissingBsc(changelog1 []searchResult, changelog2 []searchResult) []string {
	bscList1 := getBscs(changelog1)
	bscList2 := getBscs(changelog2)
	sort.Strings(bscList1)
	sort.Strings(bscList2)

	var missingBscs []string
	for _, bsc := range bscList1 {
		searchPos := sort.SearchStrings(bscList2, bsc)
		if searchPos < len(bscList2) && bscList2[searchPos] == bsc {
			// found it
		} else {
			missingBscs = append(missingBscs, bsc)
		}
	}
	return removeDuplicates(missingBscs)
}

// Extracts the BSC numbers from the search results and returns them as an array.
func getBscs(res []searchResult) []string {
	var bsc []string
	for _, v := range res {
		for _, value := range v.match {
			bsc = append(bsc, value)
		}
	}
	return bsc
}

// Scans the file for bsc, CVE and issue numbers and returns the search results.
func scanFile(pathToFile string) []searchResult {
	var regexes []*regexp.Regexp
	// creating the regexes with the regex-strings from main().
	for _, regexString := range regexStrings {
		var re, _ = regexp.Compile(regexString)
		regexes = append(regexes, re)
	}
	lines, err := scanLines(pathToFile)
	if err != nil {
		panic(err)
	}
	var searchResults []searchResult
	for i, line := range lines {
		for _, re := range regexes {
			results := re.FindAllString(line, -1)
			if len(results) > 0 {
				res := searchResult{
					line:  i + 1,
					match: results,
					text:  line}
				searchResults = append(searchResults, res)
			}
		}
	}
	return searchResults
}

// Returns the given file as an array of lines.
func scanLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

// Removes duplicates form an array.
func removeDuplicates(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
		if _, ok := m[item]; ok {
			// duplicate item
		} else {
			m[item] = true
		}
	}
	var result []string
	for item := range m {
		result = append(result, item)
	}
	return result
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
