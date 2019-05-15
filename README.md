# MiniDB

![GitHub](https://img.shields.io/github/license/ritwik310/mini-db.svg)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/ritwik310/mini-db.svg)
![Travis (.com)](https://img.shields.io/travis/com/ritwik310/mini-db.svg)

MiniDB is a very simple **in-memory key-value store**, with persistent-storage as an option. ~~Somewhat like~~ Copy of Redis. The application contains an interactive **Shell** to run DB-commands, more about it [here](#minidb-shell). And also a **Golang client** for programmatic interactions remotely, more [here](#minidb-server).

> NOTE: MiniDB is not intended to be used in a production application.

<!-- # About

Something About The App, -->

# Data Types

MiniDB stores **key-value pairs**, and supports **3 types** of values - **String**, **Integer**, **Binary**. To declare type while writing data using **MiniDB-Client** check out [this](https://github.com/ritwik310/mini-db/blob/master/client/README.md#data-type-declaration). And Type declaration in **MiniDB-Shell** has been explained [here](#commands)

> NOTE: If no data-type arguement provided in **Client** or **Shell**, MiniDB will save **String type by default**.

# Installation

Install **MiniDB-Server** using,

```shell
go get github.com/ritwik310/mini-db
```

To install the **MiniDB Go-client** run,

```shell
go get github.com/ritwik310/mini-db/client
```

# MiniDB-Server

MiniDB-Server is a **TCP server** that listens on specified port and stores data **in-memory**. Also persists data in a output file too, if option provided. You can insert or get data from the running server, using [client](#minidb-client) that writes to the TCP connection on provided address.

### Starting Server
To start the server run,
```shell
mini-db server --backup # --backup flag persists data in a output file
```

### Server Specs

By default the server starts on Port-8080. To start the server on a **custom port**,

```shell
mini-db server --port 8000
```

By default the server saves data **in-memory**. That means whenever you kill your process, you loose all data. To backup the data in a output-file add the `--backup` or `-b` flag. This takes a **snapshot** of data and persists it to filesystem. By default the data get's saved once every **5 seconds**, but you can change it using `--delay` or `-d` flag. Here's an example,

```shell
mini-db server --port 8000 --backup --delay 6 # delay in seconds (6)
```

By the way, you can even change the data **output file** using `--output` or `-o` flag,

```shell
mini-db server --port 8000 --backup --delay 1 --output ./mybackup.out
```

# MiniDB-Shell

MiniDB-Shell is an **interactive shell** for storing and manipulating key values without starting a Server. You can run [commands](#commands) in the shell and play with the DB.

### Start Shell

To start **Shell** just run,

```shell
mini-db shell
```

The Shell command also shares all the same flags that **Server** does (except for the `--port` flag), [here's](#server-specs) the spec reference. And command example,

```shell
mini-db shell --backup --delay 1 --output ./mybackup.out
```

Here's a simple demo,

<img src="https://gitlab.com/ritwik310/project-documents/raw/master/MiniDB/MiniDB-Demo-GIF-0.gif" alt="demo-gif"/>

> NOTE: Actyally GET, SET, and DELETE - these are all the commands, just 3.

### Commands

There are just three commands types to play with MiniDB data store
1. SET - Inserts data into MiniDB
2. GET - Reads the data from MiniDB and prints it
3. DELETE - Deletes data from MiniDB

> NOTE: To update data, you can use SET, it will replace the previous value.

The 2nd element of the command represents the **Key**, MiniDB-Shell doesn't count for white spaces in Key. Here's an example,

```shell
$ SET MyMsg Secret message so saving it as a binary --binary
$ GET MyMsg
Secret message so saving it as a binary
$ 
```

In the example above, **Key** is "MyMsg", and the **Value** is **binary encoded** "Secret message so saving it as a binary". The `--binary` or `--bin` flag saves it as a binary encoded string. And `GET MyMsg` returns the decoded string. You can also save **integers**, for example `SET mynum 100 --int`.

You can specify **data types** in the SET command, you need to put the type flag at the **end of the command** as shown in the example above. **By default its string** if no type provided. Here's the supported data types and corresponding flags,
1. String - `--str` or `--string`
2. Integer - `--int` or `--integer`
1. Binary - `--bin` or `--binary`

> NOTE: Everything about between the **Key** and **Type Flag** is considered to be the **Value**, this includes white spaces too.

# MiniDB Client

**MiniDB-Client** contains methods that programmatically interacts with a local or remote MiniDB-Server, and writes using TCP connection.

More about the Client [https://github.com/ritwik310/mini-db/blob/master/client/README.md](https://github.com/ritwik310/mini-db/blob/master/client/README.md)   
Read the API-Docs for Client [https://godoc.org/github.com/ritwik310/mini-db/client](https://godoc.org/github.com/ritwik310/mini-db/client)

# Happy Hacking ...