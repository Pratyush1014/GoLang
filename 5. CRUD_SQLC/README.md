# Using SQLC to generate CRUD operations

- We have multiple ways to do it
    1. DATABASE/SQL -> most basic way and very fast
    2. GORM -> go's own way of implementing less prod's code but slow and has its own syntax
    3. SQLX -> very fast
    4. SQLC -> very fast but not enough support for MySQL

- Installing SQLC
    1. Link: https://docs.sqlc.dev/en/latest/overview/install.html
    ```
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest  // using go installer
    ```
    check
    ```
    sqlc version
    sqlc help
    ```

- Navigate to Link: https://dashboard.sqlc.dev/
    1. Inside db folder create a folder for query and sqlc
    2. Create an organization and follow along
    3. Start a server using the "Create database server" button above.
    Create an auth token from the "Auth tokens" page and provide it to sqlc via the SQLC_AUTH_TOKEN environment variable.
    ```
    export SQLC_AUTH_TOKEN=sqlc_xxxxxxxx
    ```
    Copy or adapt the following sqlc.yaml config into your project (we've included the correct project ID for you):
    ```
    version: '2'
    cloud:
        project: '01HJK3FTQG89268CXPW4B4S9ND'
    sql:
        -   schema: "./db/migration" // path to your schema migration files
            queries: "./db/query"  // path to your queries
            engine: postgresql  // the database engine used
            database:
                managed: true
            gen:
                go:
                    package: "db"
                    out: "db/sqlc"
                    sql_package: "pgx/v5"
    ```
    4. Generate token looks like this, "sqlc_01HJK3H5H0ETK5QCTJST6P17FA"

- Creating CRUD operations
    1. Create 
        ```
        -- name: CreateAccount :one // name of the function 'CreateAccount' :one specifies we are retuning only one element
        INSERT INTO accounts (
          owner, balance, currency  // Unlike timestamp and id we need to insert these
        ) VALUES (
          $1, $2, $3                // three variables for the above three fields
        )
        RETURNING *;                // returns all fields of the table
        ```
        Run 
        ```
        make sqlc
        ```
        you will find go generating some files automatically, these arent supposed t be manipulated

        Next initialize a go module
        ```
        go mod init github.com/pratyush/simple_bank   // to initialize a module
        go mod tidy    // to install the dependencies of the module
        ```
    2. Read
        ```
        -- name: GetAccount :one
        SELECT * FROM accounts
        WHERE id = $1 LIMIT 1;   // Limits to only one row

        -- name: ListAccounts :many
        SELECT * FROM accounts
        ORDER BY id
        LIMIT $1
        OFFSET $2;              // We do not fetch all records at once rather we skip 2 records sort of pagination
        ```

    3. Similarly refer to account.sql to see update and delete operations.