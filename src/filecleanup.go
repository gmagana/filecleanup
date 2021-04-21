package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	fmt.Println("filecleanup v1.0.0 - https://github.com/gmagana/filecleanup")
	fmt.Println("© 2021 - Gabriel Magana-Gonzalez - gmagana@gmail.com")
	fmt.Println()
	//fmt.Println(fmt.Sprintf("Args (Count: %d):", len(os.Args[1:])))
	//fmt.Println(os.Args[1:])

	var liveRun bool
	var listAllFiles bool
	var caseInsensitiveSort bool
	var reverseOrder bool
	var filesToKeep int

	flag.BoolVar(&liveRun,
		"live-run",
		false,
		"If specified, files are deleted. If not specified, a dry run takes place (no files are deleted)")

	flag.BoolVar(&listAllFiles,
		"list-all-files",
		false,
		"Show list of all files in order in which they will be processed")

	flag.BoolVar(&caseInsensitiveSort,
		"order-case-insensitive",
		false,
		"Sorts filenames in case insensitive order")

	flag.BoolVar(&reverseOrder,
		"order-reverse",
		false,
		"Sorts filenames in reverse order")

	flag.IntVar(&filesToKeep,
		"files-to-keep",
		0,
		"(REQUIRED) The number of files you intend to exist after the deletion takes place")

	flag.Parse()

	AssureRequiredFlagsPresent([]string{"files-to-keep"})

	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "ERROR: Please specify one file pattern filter (Arguments: %s)\n", flag.Args())
		os.Exit(2) // the same exit code flag.Parse uses
	}

	targetPattern := flag.Args()[0]

	fmt.Println("\tOPTIONS:")
	fmt.Println(fmt.Sprintf("\tDRY RUN: %t", !liveRun))
	fmt.Println(fmt.Sprintf("\tLIST ALL FILES: %t", listAllFiles))
	fmt.Println(fmt.Sprintf("\tCASE INSENSITIVE ORDER: %t", caseInsensitiveSort))
	fmt.Println(fmt.Sprintf("\tREVERSE ORDER: %t", reverseOrder))
	fmt.Println(fmt.Sprintf("\tTARGET PATTERN: %s", targetPattern))
	fmt.Println(fmt.Sprintf("\tFILES TO KEEP: %d", filesToKeep))

	files, err := filepath.Glob(targetPattern)

	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: %s", err))
		os.Exit(2)
	}

	// Sort existing filenames depending on specified flags
	switch {
	case caseInsensitiveSort && reverseOrder:
		sort.Slice(files,
			func(i, j int) bool {
				return strings.ToLower(files[i]) > strings.ToLower(files[j])
			})
	case !caseInsensitiveSort && reverseOrder:
		sort.Slice(files,
			func(i, j int) bool {
				return files[i] > files[j]
			})
	case caseInsensitiveSort && !reverseOrder:
		sort.Slice(files,
			func(i, j int) bool {
				return strings.ToLower(files[i]) < strings.ToLower(files[j])
			})
	case !caseInsensitiveSort && !reverseOrder:
		sort.Slice(files,
			func(i, j int) bool {
				return files[i] < files[j]
			})
	default:
		panic("Not implemented")
	}

	if listAllFiles {
		fmt.Println()
		fmt.Println("All files (sorted):")
		for _, file := range files {
			fmt.Println(file)
		}
		fmt.Println()
	}

	fileCountToDelete := len(files) - filesToKeep
	var filesToDelete []string = nil
	if fileCountToDelete >= 1 {
		filesToDelete = files[0:fileCountToDelete]
	} else {
		fileCountToDelete = 0
	}

	if !liveRun {
		fmt.Println("*** DRY RUN: No files will be deleted (run with --live-run parameter to delete files) ***")
	}
	fmt.Println(fmt.Sprintf("Files found: %d - Files to delete: %d", len(files), fileCountToDelete))
	for _, file := range filesToDelete {
		fmt.Println(fmt.Sprintf("DELETING FILE: %s", file))
		if liveRun {
			err := os.Remove(file)

			if err != nil {
				fmt.Println(fmt.Sprintf("ERROR DELETING FILE: %s", err))
			}
		}
	}
	if !liveRun {
		fmt.Println("*** DRY RUN: No files deleted ***")
	}
	fmt.Println()

	fmt.Println("Done.")
}

func AssureRequiredFlagsPresent(required []string) {
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			fmt.Fprintf(os.Stderr, "ERROR: Missing required --%s argument\n", req)
			os.Exit(2) // the same exit code flag.Parse uses
		}
	}
}
