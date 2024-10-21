# cpdiff

## Installation

### From source

```sh
go install github.com/ChrisVilches/cpdiff@latest
```

### Arch Linux (AUR)

```sh
yay -S cpdiff
```

## Development

Linting and formatting:

```sh
go fmt ./...
revive --formatter stylish ./...
```

In order to run tests, first build the project, then set the environment variable below to the path of the executable:

```sh
go build
INTEGRATION_TEST_EXECUTABLE=/path/cpdiff go test ./...
```
