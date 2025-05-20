# QCache - In-Memory Key-Value Store

[![Go Report Card](https://goreportcard.com/badge/github.com/ayuuuuu0-0/qcache)](https://goreportcard.com/report/github.com/ayuuuuu0-0/qcache)

QCache is a high-performance, in-memory key-value store built in Go. It leverages goroutines for concurrent request handling and implements the Redis RESP protocol for client-server communication.

## Features

- **High Performance**: Optimized for fast read/write operations with minimal latency
- **Concurrent Processing**: Uses Go's goroutines for handling multiple client connections simultaneously
- **Redis Protocol Support**: Implements the RESP (Redis Serialization Protocol) for compatibility
- **Thread-safe Operations**: Ensures data consistency with proper synchronization mechanisms
- **Simple Client Interface**: Easy-to-use client library for seamless integration

## Getting Started

### Prerequisites

- Go 1.18 or higher
- Git

### Installation

```bash
# Clone the repository
git clone https://github.com/ayuuuuu0-0/qcache.git
cd qcache

# Build the binary
make build
```

### Running the Server

```bash
make run
# OR
./bin/qcache --listenAddr :5001
```

## Usage

### Connecting with a Client

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/ayuuuuu0-0/qcache/client"
)

func main() {
    // Connect to QCache server
    c, err := client.New("localhost:5001")
    if err != nil {
        log.Fatal(err)
    }
    
    // Store a value
    if err := c.Set(context.TODO(), "mykey", "myvalue"); err != nil {
        log.Fatal(err)
    }
    
    // Retrieve a value
    val, err := c.Get(context.TODO(), "mykey")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Retrieved value:", val)
}
```

### Using Telnet for Testing

You can also connect to QCache using telnet to test basic functionality:

```
telnet localhost 5001
*3
$3
SET
$5
hello
$5
world
*2
$3
GET
$5
hello
```

## Architecture

QCache is built with a focus on concurrent processing and low-latency operations:

```
┌──────────────┐         ┌──────────────┐
│              │         │              │
│    Client    │◄───────►│   Server     │
│              │  RESP   │              │
└──────────────┘         └───────┬──────┘
                                 │
                          ┌──────▼──────┐
                          │             │
                          │  Handler    │
                          │             │
                          └──────┬──────┘
                                 │
                          ┌──────▼──────┐
                          │             │
                          │ In-Memory   │
                          │   Store     │
                          │             │
                          └─────────────┘
```

The server uses a message-passing architecture with channels to handle commands:

```go
func(s *Server) loop() {
    for {
        select {
        case msg := <- s.msgCh:
            if err := s.handleMessage(msg); err != nil {
                slog.Error("raw message error", "err", err)
            }
        case <- s.quitCh:
            return 
        case peer := <- s.addPeerCh:
            s.peers[peer] = true
        }
    }
}
```

## Development

### Project Structure

```
qcache/
├── bin/              # Compiled binary
├── client/           # Client library package
├── main.go           # Main server entry point
├── peer.go           # Connection handling
├── resp.go           # RESP protocol implementation
└── Makefile          # Build commands
```

### Building and Running

```bash
# Build the binary
make build

# Run the server
make run
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgements

- Inspired by Redis and BoltDB
- Built with Go's powerful concurrency primitives

---

[GitHub](https://github.com/ayuuuuu0-0/qcache) | [Report Issues](https://github.com/ayuuuuu0-0/qcache/issues)
