# Touch Fake

![GitHub](https://img.shields.io/github/license/NalbertLeal/touch-fake)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/NalbertLeal/touch-fake)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/NalbertLeal/touch-fake)

## Description

_touch-fake_ is a Go tool that allows you to easily create fake files with a specific size in bytes. This can be useful for performance testing, simulations, and other scenarios where you need a file of a certain size without the need for real content. The name touch fake was based on the command _touch_ used on linux to create a file.

## How to Use

### Prerequisites

Make sure you have Go installed on your machine. For more information on how to install Go, visit [https://golang.org/doc/install](https://golang.org/doc/install).

### Installation

```bash
go get -u github.com/NalbertLeal/touch-fake
```

### Basic Usage

This will create a file named `fake_file.txt` with a size of 1024 bytes:

```bash
touch-fake fake_file.txt 1024 
```

And this will create a file named `fake_file.txt` with a size of 1024 bytes that overwrite the file (if already exist):

```bash
touch-fake fake_file.txt 1024 -overwrite
```

### Options

- `-overwrite`: Overwrite the output file if it already exists (optional).

## Contribution

If you would like to contribute improvements, bug fixes, or new features, feel free to open an issue or submit a pull request. Your contribution is welcome!

## License

This project is licensed under the [GNU LESSER GENERAL PUBLIC LICENSE](https://choosealicense.com/licenses/lgpl-3.0/).

---
