package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zenizh/go-capturer"
)

// ----------------------------------------------------------------------------
//  Examples (Tests for golden-cases)
// ----------------------------------------------------------------------------

func ExampleGetInfo() {
	namePkg := "github.com/KEINOS/Hello-Cobra"

	// The output format is as follows:
	//   - Hello-Cobra
	//     1. Gravity:     11
	//     2. Package Name: github.com/KEINOS/Hello-Cobra
	//     3. URL:          https://github.com/KEINOS/Hello-Cobra
	//     4. Stars:        11
	//     5. Forks:        2
	//     6. Folllows:     2
	//     7. ImportedBy:   0
	infoPkg, err := GetInfo(namePkg)
	if err != nil {
		log.Fatal(err)
	}

	contains := true

	for _, contain := range []string{
		"Hello-Cobra",
		"Gravity:",
		"URL:",
		"https://github.com/KEINOS/Hello-Cobra",
		"Stars:",
		"Forks:",
		"Folllows:",
		"ImportedBy:",
	} {
		contains = strings.Contains(infoPkg, contain) && contains
	}

	fmt.Printf("The output contains all: %v", contains)

	// Output:
	// The output contains all: true
}

func ExampleSprintStringMap() {
	input := map[string]interface{}{
		"ten":      10,
		"thousand": 1000,
		"eleven":   11,
		"one":      "uno",
	}

	fmt.Println(SprintStringMap(input))

	// Output:
	// eleven:   11
	// one:      uno
	// ten:      10
	// thousand: 1000
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func Test_main_golden(t *testing.T) {
	// Backup and defer restore
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// Mock os.Args
	os.Args = []string{
		t.Name(),
		"github.com/KEINOS/Hello-Cobra",
	}

	out := capturer.CaptureOutput(func() {
		main()
	})

	for _, contain := range []string{
		"Hello-Cobra",
		"Gravity:",
		"URL:",
		"https://github.com/KEINOS/Hello-Cobra",
		"Stars:",
		"Forks:",
		"Folllows:",
		"ImportedBy:",
	} {
		require.Contains(t, out, contain)
	}
}

func Test_main_help(t *testing.T) {
	// Backup and defer restore
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// Mock os.Args
	os.Args = []string{
		t.Name(),
	}

	out := capturer.CaptureOutput(func() {
		main()
	})

	assert.Contains(t, out, "help me")
}

func TestExitOnError(t *testing.T) {
	// Backup and defer restore
	oldLogFatal := LogFatal
	defer func() {
		LogFatal = oldLogFatal
	}()

	var (
		errDummy    = errors.New("forced error")
		errCaptured error
		ok          bool
	)

	// Mock LogFatal and capture the given arg
	LogFatal = func(v ...interface{}) {
		errCaptured, ok = v[0].(error)

		require.True(t, ok)
	}

	// Test
	ExitOnError(errDummy)

	expect := "forced error"
	actual := errCaptured.Error()

	assert.Equal(t, expect, actual)

}

func TestGetInfo_bad_package_name(t *testing.T) {
	namePkg := "github.com/KEINOS/undefined"
	expectErr := "failed to get 'imported by' information"

	out, err := GetInfo(namePkg)

	require.Error(t, err)
	assert.Contains(t, err.Error(), expectErr)
	assert.Empty(t, out, "it should be empty on error")
}
