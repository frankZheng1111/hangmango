# hangmango
Just for practice：A simple command line hangman game based on golang

## Install dependency

- need [glide](https://github.com/Masterminds/glide)
  ```
  glide install
  ```

## Run unit test

- run all unit test
  ```
  GO_ENV=test go test -coverprofile coverage.out ./...
  ```

- get total coverage
  ```
  go tool cover -func=coverage.out
  ```
