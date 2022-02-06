# GoStars - Get the Attraction Gravity of a Go Package

`gostars` retuns the attraction graivity of a Go package in GitHub.

The aim of `gostars` is to provide as much information as possible to those who are struggling to decide which package to use, so that they can make an informed decision.

```bash
# Usage
gostars <package name> [...<package name>]
```

```shellsession
$ # Sample
$ gostars github.com/daviddengcn/go-colortext
- go-colortext
  1. Gravity:      583
  2. Package Name: github.com/daviddengcn/go-colortext
  3. URL:          https://github.com/daviddengcn/go-colortext
  4. Stars:        209
  5. Forks:        20
  6. Folllows:     9
  7. ImportedBy:   544
```

```shellsession
$ # Sample
$ gostars github.com/goccy/go-json github.com/json-iterator/go
- go-json
  1. Gravity:      1331
  2. Package Name: github.com/goccy/go-json
  3. URL:          https://github.com/goccy/go-json
  4. Stars:        1321
  5. Forks:        48
  6. Folllows:     17
  7. ImportedBy:   160
- go
  1. Gravity:      12470
  2. Package Name: github.com/json-iterator/go
  3. URL:          https://github.com/json-iterator/go
  4. Stars:        10426
  5. Forks:        850
  6. Folllows:     238
  7. ImportedBy:   6785
```

## About "Gravity"

The element name "Gravity" represents the suction force of the Go package.

- The current basic formula for measuring attractiveness:

  ```go
  weight := math.Sqrt(
      (numStars * numStars) +
      (numFolows * numFolows) +
      (numForks * numForks) +
      (numImports * numImports),
  )
  ```

## Install

```bash
# macOS/Linux (x86_64/AMD64/ARM64/M1)
brew install KEINOS/apps/gostars
```

- For Windows, ARM5/6/7, etc. see the [releases page](https://github.com/KEINOS/gostars/releases).

## License

- License: [MIT](https://github.com/KEINOS/gostars/LICENSE.txt)
- Copyright: (c) KEINOS and [the GoStars contributors](https://github.com/KEINOS/gostars/graphs/contributors).

## ToDo

- Add more elements to the formula to measure weight.
  - [ ] Add [Go Report Card](https://goreportcard.com/)'s score.
  - [ ] Add code coverage rate.
  - [ ] Add "Code frequency" or "Pulse".