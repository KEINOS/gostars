package gostars_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/KEINOS/gostars/gostars"
	"github.com/PuerkitoBio/goquery"
)

func ExampleGetAttractionGravity() {
	{
		forks := 1
		likes := 10
		followers := 0
		importedBy := 100

		gravity := gostars.GetAttractionGravity(forks, likes, followers, importedBy)

		fmt.Println("Light Star:", gravity)
	}
	{
		forks := 10
		likes := 100
		followers := 10
		importedBy := 1000

		gravity := gostars.GetAttractionGravity(forks, likes, followers, importedBy)

		fmt.Println("Heaby Star:", gravity)
	}

	// Output:
	// Light Star: 100
	// Heaby Star: 1005
}

func ExampleGetContentURL() {
	rawContent, err := gostars.GetContentURL("https://github.com/KEINOS")
	if err != nil {
		log.Fatal(err)
	}

	source := string(rawContent)

	fmt.Println(strings.Contains(source, "Profile of KEINOS"))

	// Output: true
}

func ExampleGetURLGitHub() {
	urlSite := "https://joe-bot.net/"

	// Get the actual URL of the GitHub repository from the mapping
	urlGitHub := gostars.GetURLGitHub(urlSite)

	fmt.Println(urlGitHub)

	// Output: https://github.com/go-joe/joe
}

func ExampleHash256() {
	out := gostars.Hash256([]byte("sample"))

	fmt.Println(out)
	// Output: af2bdbe1aa9b6ec1e2ade1d694f41fc71a831d0268e9891562113d8a62add1bf
}

func ExampleNewPkgInfo() {
	// Get package info from "pkg.go.dev"
	pkgInfo, err := gostars.NewPkgInfo("github.com/KEINOS/go-utiles/util")
	if err != nil {
		log.Fatal(err)
	}

	// Get number of packages that uses this package
	if pkgInfo.ImportedBy > 3 {
		fmt.Println("This package has been used by 3 or more packages.")
	}

	fmt.Println("The URL of the repository:", pkgInfo.Repository)

	// Output:
	// This package has been used by 3 or more packages.
	// The URL of the repository: https://github.com/KEINOS/go-utiles
}

func ExampleNewRepoInfo() {
	repoInfo, err := gostars.NewRepoInfo("https://github.com/KEINOS/dev-go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name repo:", repoInfo.Name)
	fmt.Println("Name owner:", repoInfo.Owner)

	// Output:
	// Name repo: dev-go
	// Name owner: KEINOS
}

func ExampleNewQuery() {
	html := `
<body>
<ul>
<li><a href="link_to_foo">foo</a></li>
<li><a href="link_to_bar">bar</a></li>
</ul>
</body>
	`

	goQuery, err := gostars.NewQuery([]byte(html))
	if err != nil {
		log.Fatal(err)
	}

	foundLinks := make([]string, 0) // Var to store found links

	// Callback function to store the link if the selection s contains an href attribute.
	setHref := func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			log.Fatal("fail to get attribute's value for the first element in the selection")
		}

		foundLinks = append(foundLinks, href)
	}

	// Query to find "body.li.a" element and iterate with `setHref` function.
	goQuery.Find("body li > a:first-child").Each(setHref)

	// Print the result
	for i, link := range foundLinks {
		fmt.Println("#", i, "LINK:", link)
	}

	// Output:
	// # 0 LINK: link_to_foo
	// # 1 LINK: link_to_bar
}

func ExampleNewURLInfo() {
	urlInfo, err := gostars.NewURLInfo("https://github.com/KEINOS/gostars")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scheme:", urlInfo.Scheme)
	fmt.Println("Host:", urlInfo.Host)
	fmt.Println("Is GitHub repo:", urlInfo.IsRepoGitHub())
	fmt.Println("Stringer:", urlInfo)

	// Output:
	// Scheme: https
	// Host: github.com
	// Is GitHub repo: true
	// Stringer: https://github.com/KEINOS/gostars
}

func ExampleParseMarkdownToHTML() {
	markdown := `
# Hello
Hello, world!
	`

	out := gostars.ParseMarkdownToHTML([]byte(markdown))

	fmt.Println(out)

	// Output:
	// <body><h1>Hello</h1>
	//
	// <p>Hello, world!</p>
	// </body>
}

func ExamplePrettyFormatJSON() {
	type myStruct struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}

	myObj := myStruct{
		Foo: "hoge",
		Bar: "fuga",
	}

	result, err := gostars.PrettyFormatJSON(myObj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	// Output:
	// {
	//   "foo": "hoge",
	//   "bar": "fuga"
	// }
}
