# MiniDB Client

![GitHub](https://img.shields.io/github/license/ritwik310/mini-db.svg)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/ritwik310/mini-db.svg)
![Travis (.com)](https://img.shields.io/travis/com/ritwik310/mini-db.svg)

**MiniDB-Client** contains methods that programmatically interacts with the [**MiniDB-Server**](https://github.com/ritwik310/mini-db/#mini-db).

# Quick start

### Installation

First, install MiniDB using and MiniDB Golang-client

```shell
go get github.com/ritwik310/mini-db
go get github.com/ritwik310/mini-db/client
```

### Start Server

Start a TCP-Server using MiniDB-CLI, details about starting the **MiniDB-Server** and other Server-related options [here](https://github.com/ritwik310/mini-db)

```shell
mini-db server --backup # --backup persists the data in the filesystem
```

### Interacting with Server

Here's a pretty straightforward **code-example** that interacts with the Server using **MiniDB-Client**

```go
package main

import (
    "fmt"
    "github.com/ritwik310/mini-db/client"
)

func main() {
    // Create a client instance (mdb)
    mdb := client.New("tcp", "localhost:8000") // By default the MiniDB-Server listens on Port-8000

    // Communicating to the Server
    resp0, err := mdb.Set("myname", "Ritwik Saha", "str") // "myname" => key, "Ritwik Saha" => value, "str" => data-type
    resp1, err := mdb.Get("myname")
    resp2, err := mdb.Delete("myname")
    resp3, err := mdb.Get("myname")

    if err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Printf("resp0: %+v\n", resp0) // resp0["status"] == 200
    fmt.Printf("resp1: %+v\n", resp1) // resp1["data"] == "Ritwik Saha"
    fmt.Printf("resp2: %+v\n", resp2) // resp2["status"] == 200
    fmt.Printf("resp3: %+v\n", resp3) // resp3["error"] != nil && resp3["status"] == 400
}

```

# Documentation

Read the API-Docs [godoc.org/github.com/ritwik310/mini-db/client](https://godoc.org/github.com/ritwik310/mini-db/client)
