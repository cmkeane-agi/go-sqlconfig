# go-sqlconfig

Import and usage
---------------

Import the module:

- import "github.com/cmkeane-agi/go-sqlconfig"

Example:

```go
package main

import (
    "fmt"
    "github.com/cmkeane-agi/go-sqlconfig"
)

func main() {
    store, err := sqlconfig.LoadSQL("example.sql")
    if err != nil {
        panic(err)
    }
    fmt.Println(store.Must("one"))
}
```
## go-sqlconfig: A Simple SQL File Loader for Go Applications

go-sqlconfig is a lightweight Go package that simplifies the management and loading of SQL queries from external files. It allows developers to organize their SQL queries in a structured manner, making it easier to maintain and update database interactions within Go applications.

### Features
- Load SQL queries from external `.sql` files.
- Organize queries using named sections for easy retrieval.
- Simple API for integrating SQL queries into Go code.
- Supports parameterized queries for enhanced security.
- Easy to use and integrate into existing Go projects.

### Installation
To install go-sqlconfig, use the following command:
```bash
go get github.com/cmkeane-agi/go-sqlconfig
```

### Usage
1. Create a `.sql` file with your SQL queries, using the `-- name: <query_name>` format to define named sections.
2. Use the go-sqlconfig package to load and retrieve your SQL queries in your Go application.
```go
package main
import (
    "fmt"
    "log"

    "github.com/cmkeane-agi/go-sqlconfig"
)

func main() {
    // Load SQL queries from file
    sqlLoader, err := sqlconfig.NewLoader("path/to/your/queries.sql")
    if err != nil {
        log.Fatalf("Failed to load SQL file: %v", err)
    }

    // Retrieve a specific query by name
    query, err := sqlLoader.GetQuery("login.current")
    if err != nil {
        log.Fatalf("Failed to get query: %v", err)
    }

    fmt.Println("SQL Query:", query)
}
``` 

### Contributing
Contributions are welcome! Please feel free to submit issues or pull requests on the GitHub repository.