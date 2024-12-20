# cpdiff

File difference tool for competitive programming.

## Installation

```sh
go install github.com/ChrisVilches/cpdiff@latest
```

### Arch Linux (AUR)

```sh
yay -S cpdiff
```

## Usage

Comparing standard input (e.g. some program's output redirected to `cpdiff`) against a file (e.g. the correct answer):

```sh
./my-program < in-file | cpdiff ans-file
```

Or compare two files:

```sh
cpdiff file1 file2
```

Help (see all options):

```sh
cpdiff --help
```

## Screenshots

Comparing each number individually. Numbers are allowed to have an error:

![numbers](./screenshots/numbers.png)

Comparing strings. Each character is compared individually:

![heart-strings](./screenshots/heart-strings.png)

If you want to compare binary strings (or digits), you can compare them character by character instead of comparing their numeric value:

![binary-strings](./screenshots/binary-strings.png)

Big numbers are supported. Numbers can have arbitrary precision or amount of digits:

![big-numbers](./screenshots/big-numbers.png)

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

Verify installation command works:

```sh
IMG=cpdiff-install-test

docker build -t $IMG .
docker run --rm $IMG
```

Clean up manually afterwards (i.e. remove images).

Generate screenshots by first installing [termshot](https://github.com/homeport/termshot) and then executing the following command:

```sh
bash ./screenshots/gen.sh path_to_termshot_executable path_to_cpdiff_executable
```
