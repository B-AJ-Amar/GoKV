
# GoKV

GoKV is a simple, in-memory key–value store implemented in Go, designed as a lightweight, Redis-compatible alternative for learning, prototyping, or embedded use.

## Features

- **RESP Protocol Support**: Parses and responds using the Redis Serialization Protocol (RESP), supporting most basic Redis commands.
- **Commands Supported**:
	- `SET` / `GET`: Store and retrieve string values.
	- `DEL`: Delete one or more keys.
	- `EXISTS`: Check if keys exist.
	- `INCR` / `INCRBY`: Atomic integer increment operations.
	- `SELECT`: Switch between logical databases (multi-DB support).
	- `PING`: Health check.
	- `HELLO`: RESP version negotiation.
- **SETX Command Extensions**:
	- `NX` / `XX`: Set if not exists / set if exists.
	- `EX` / `PX`: Expiration in seconds or milliseconds.
	- `KEEPTTL`: Retain existing TTL.
	- `GET`: Return old value on set.
- **Expiration**: Key expiration with millisecond precision.
- **Multiple Databases**: Supports multiple logical databases, switchable via `SELECT`.
- **Centralized Error Handling**: All errors are defined in a single location for maintainability.


## Usage

1. **Run the server:**
	 ```sh
	 go run ./cmd/gokv/main.go
	 ```
2. **Connect with redis-cli or any RESP-compatible client:**
	 ```sh
	 redis-cli -p 6379
	 ```

## TODO
- [ ] Add internal debug logs for easier troubleshooting.
- [ ] Add queue support to the store (for future data structures).
- [ ] Add mutexes for thread-safe `Set` and `Setx` operations.
- [ ] Add more advanced Redis commands (e.g., `MGET`, `MSET`, `EXPIRE`, `TTL`).
- [ ] Improve error messages and RESP compliance.
- [ ] Add authentication and ACL support.
- [ ] Add more comprehensive tests and benchmarks.


