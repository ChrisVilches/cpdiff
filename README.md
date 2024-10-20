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

In order to run tests, first build the project, then set the environment variable below to the path of the executable:

```sh
go build
INTEGRATION_TEST_EXECUTABLE=/path/cpdiff go test ./...
```
