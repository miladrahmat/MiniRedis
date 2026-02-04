<div align=center>
  <h1>MiniRedis</h1>
  <p>A Redis-compatible server built with Go</p>
</div>

## Content

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Supported commands](#supported-commands)



## Introduction

This project is a lightweight, high-performance Redis-compatible server built from scratch using Go. The server implements the [Redis Serialization Protocol (RESP)](https://redis.io/docs/latest/develop/reference/protocol-spec/) and supports concurrent client connections using Go's native concurrency primitives. It utilizes `sync.RWMutex` to ensure thread-safety and prevent data races during simultaneous read/write operations.

## Features

- **RESP Protocol Support:** Custom built parser to handle Redis commands.
- **Multiple client connections:** Leverages Go's goroutines to handle multiple simultaneous client connections efficiently.
- **Redis CLI Compatible:** Works with the standard `redis-cli` or any Redis-compatible client library.
- **In-Memory Storage:** Fast data access and manipulation.

## Installation

1. Clone the repository in the directory of your choice:

```bash
git clone https://github.com/miladrahmat/MiniRedis.git
cd MiniRedis
```

2. Build the project:

```bash
go build ./cmd/mini-redis/mini-redis.go
```

3. Start the server:

```bash
./mini-redis
```

## Usage

Once the server is running you can connect to it using the standard `redis-cli`:

```bash
redis-cli -p 6379
```

#### Supported commands:

- `PING` -> Responds with `PONG`
- `SET <key> <value>` -> Store a key-value pair
- `GET <key>` -> Retreive a value by it's key
- `HSET <key> <field> <value>` -> Sets the specified field in the hash stored at `<key>` to `<value>`
- `HGET <key> <field>` -> Retreive the value assiociated with `<field>` in the hash stored at `<key>`
- `HGETALL <key>` -> Retreive all fields and their values of the hash stored at `<key>`

## Future features

This project is an ongoing development project for me to learn Go and understand databases on a deeper level. While the core RESP parser and concurrency model are functional, the following features are planned to be developed:

1. Data persistence (AOF)

To prevent data loss on restart or in case of a crash, I plan to implement an **Append Only File (AOF)** where every operation will be logged.

2. Expanded command set

Adding support for more complex data structures like **Lists** (`LPUSH`, `LPOP`) and **Sets** (`SADD`, `SMEMBERS`)
