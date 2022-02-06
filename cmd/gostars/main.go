package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/KEINOS/gostars/gostars"
)

// LogFatal is a copy of log.Fatal to ease mock during test.
var LogFatal = log.Fatal

// OsExit is a copy of os.Exit to ease mock during test.
var OsExit = os.Exit

func main() {
	if len(os.Args) > 1 {
		for _, namePackage := range os.Args[1:] {
			info, err := GetInfo(namePackage)
			ExitOnError(err)

			fmt.Println(info)
		}
	} else {
		PrintHelp()
	}
}

func GetInfo(namePkg string) (string, error) {
	var (
		pkgInfo  *gostars.PkgInfo
		repoInfo *gostars.RepoInfo
		err      error
	)

	// Get package and repository info
	if pkgInfo, err = gostars.NewPkgInfo(namePkg); err == nil {
		repoInfo, err = gostars.NewRepoInfo(pkgInfo.Repository)
	}

	if err != nil {
		return "", err
	}

	// Get dimention values from both package and repository info
	stars := repoInfo.Stars
	forks := repoInfo.Forks
	follow := repoInfo.Followers // equivalent to watching
	importedBy := pkgInfo.ImportedBy

	// Calculate force/attractiveness from the values of dimentions
	gravity := gostars.GetAttractionGravity(
		stars, forks, follow, importedBy,
	)

	indent := "  "

	result := fmt.Sprintln("-", repoInfo.Name)

	items := map[string]interface{}{
		indent + "1. Gravity":      gravity,
		indent + "2. Package Name": pkgInfo.Name,
		indent + "3. URL":          pkgInfo.Repository,
		indent + "4. Stars":        stars,
		indent + "5. Forks":        forks,
		indent + "6. Folllows":     follow,
		indent + "7. ImportedBy":   importedBy,
	}

	result += SprintStringMap(items)

	return result, nil
}

func PrintHelp() {
	fmt.Println("help me")
}

func ExitOnError(err error) {
	if err != nil {
		LogFatal(err)
	}
}

func SprintStringMap(input map[string]interface{}) string {
	keys := make([]string, 0)
	maxLen := 0

	for key := range input {
		keys = append(keys, key)

		if lenKey := len(key); lenKey > maxLen {
			maxLen = lenKey
		}
	}

	sort.Strings(keys)

	padding := strings.Repeat(" ", maxLen)
	result := ""

	for _, key := range keys {
		col1st := key + ":" + padding
		col2nd := input[key]

		result += fmt.Sprintln(col1st[0:maxLen+1], col2nd)
	}

	return "  " + strings.TrimSpace(result)
}
