# KKCOMPANY

## Application Design

The `kkcompany` application is designed to provide a suite of file utilities including line counting, checksum computation, and version information. The application is built using the Cobra library for command-line interface management.

## 3rd-Party Libraries Used

- [Cobra](https://github.com/spf13/cobra): Used for creating the command-line interface.
- [Go](https://golang.org/): The programming language used for developing the application.

## How to Build the Project

### Prerequisites

- Go 1.16 or higher

### Build

To build the project, run the following command:

```sh
go build -o futil main.go
```

### Add to binary

```
 sudo mv futil /usr/local/bin/futil
```

## How to Download/Install release version

### Prerequisites

- Go 1.16 or higher

### Download

Download suitable binary file and install.sh file from:

```
https://github.com/vuongthanh148/kkcompany/releases
```

### Install

```
 chmod +x install.sh
```

```
./install.sh
```

## How to Run the app

### Run

To run the built executable, use the following command:

```
futil <command> [flags]
```

## Features Not Yet Implemented

1. Multi-threaded processing for large files

## Known Issues

1. Limited error handling for unsupported file types
