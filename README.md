# cpdiff

## Installation

### From source

```sh
git clone https://github.com/ChrisVilches/cpdiff
cd cpdiff
go install
```

### Arch Linux (AUR)

```sh
yay -S cpdiff
```

## Development

Linting

```sh
revive --formatter stylish ./...
```

For unit and integration tests, build the project, then set its path to the environment variable:

```sh
go build
INTEGRATION_TEST_EXECUTABLE=/path/cpdiff go test ./...
```
