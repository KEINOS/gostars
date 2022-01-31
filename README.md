# GoStars - Get the Attraction Gravity of a Go Package

`gostars` retuns the attraction graivity of a Go package in GitHub.

```bash
gostars <package name> [...<package name>]

# Returns the weight of the package based on the number of stars, forks, followers and imports.
```

```shellsession
$ gostars github.com/KEINOS/myapp1
10
$ gostars github.com/KEINOS/myapp2
5
```

- The current basic formula for measuring attractiveness:
  ```go
  weight := math.Sqrt(
      (numStars * numStars) +
      (numFolows * numFolows) +
      (numForks * numForks) +
      (numImports * numImports),
  )
  ```

- The aim of `gostars` is to provide as much information as possible to those who are struggling to decide which package to use, so that they can make an informed decision.

## License

- License: [MIT](https://github.com/KEINOS/gostars/LICENSE.txt)
- Copyright: (c) [Awesome-Go](https://github.com/avelino/awesome-go/graphs/contributors) and [The GoStars](https://github.com/KEINOS/gostars/graphs/contributors) contributors.

## ToDo

- Add more elements to the formula to measure weight.
  - [ ] Add [Go Report Card](https://goreportcard.com/)'s score.
  - [ ] Add code coverage rate.
  - [ ] Add "Code frequency" or "Pulse".