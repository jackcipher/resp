# Redis RESP Converter

## Overview

A simple Golang CLI tool to convert Redis commands to RESP (REdis Serialization Protocol) format.

## What is RESP?

RESP (REdis Serialization Protocol) is the protocol used by Redis to communicate between clients and servers. This tool helps developers and Redis enthusiasts understand and visualize the underlying protocol structure.

## Features

- Command-line RESP conversion
- Interactive mode
- Support for single and multiple word commands
- Easy to use and integrate

## Installation

```bash
# Install directly from GitHub
go install github.com/jackcipher/resp@latest

# Or clone and install locally
git clone https://github.com/jackcipher/resp.git
cd resp
go install
```

## Usage

### One-shot Mode
```bash
# Convert a single command
resp SET key value
```

### Interactive Mode
```bash
# Start interactive mode
resp

# Enter commands
> SET mykey hello
*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\nhello\r\n

# Exit with 'exit' or 'quit'
> exit
```

## Examples

### Sample Conversions

| Command         | RESP Output                                           |
|----------------|-------------------------------------------------------|
| `SET key val`  | `*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$3\r\nval\r\n`        |
| `GET key`      | `*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n`                    |
| `HSET hash f v`| `*4\r\n$4\r\nHSET\r\n$4\r\nhash\r\n$1\r\nf\r\n$1\r\nv\r\n` |



## License

Distributed under the MIT License. See `LICENSE` for more information.

