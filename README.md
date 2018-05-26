# hangmango
Just for practice：A simple command line hangman game based on golang

## Run unit test

- run all unit test
  ```
  go test -coverprofile coverage.out ./...
  ```

- get total coverage
  ```
  go tool cover -func=coverage.out
  ```
