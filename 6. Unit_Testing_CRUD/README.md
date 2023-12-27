# UNIT Testing the CRUD operations using random data.

## Setting up db driver
- Why? Becoz we just have an interface to see the db but to connect to db and communicate with it we need drivers.
- We will make use of library for postgresql
    Link: github.com/lib/pq
    Navigate to install and install it
    ```
    go get github.com/lib/pq
    ```
- After installation's complete you can find the import in go.mod file with an indirect flag which means this dependency is added but hasn't been imported and used anywhere

## Creating the entry point for all test scripts
- Create a file named main_test.go it serves as the entry point for all test scripts reciding in the db pckg.

```
package db

import ("database/sql"                                      // importing dependencies
    "log"
    "os"
    "testing"

    _ "github.com/lib/pq"                                    // importing driver, go formatter usually will remove this to keep it use _
    )

const (                                                     // defining dbDriver and Source for connection
    dbDriver = "postgres"
    dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries                                    // declaring the Queries object

func TestMain (m *testing.M) {                              // Testing out the connection
    conn, err := sql.Open(dbDriver, dbSource)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    testQueries = New(conn)

    os.Exit(m.Run())
}
```
- Voila now the setups done you can proceed with writing test scripts for functions
## Testing the createAccount function

- createAccount function recides in account.sql.go, so we need to define the test scripts in the same directory
  Go follows a convention i.e. name your test script as 'account_test.sql.go'.

  CreateAccount function
  ```
  func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) { // This takes in Queries object  which needs db connection
	row := q.db.QueryRow(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
    }
  ```

  account.sql.go
  ```
  package db                                // its the same package as in your db

  import ("testing"                          // library for testing
    "context"
    "github.com/stretchr/testify/require"   // This is a utility lib install it using go get ....
    )

  func TestCreateAccount(t *testing.T) {    // This is the test function, this function always takes in a testing object
        arg:= CreateAccountParams{
            Owner: "tom",
            Balance: 100,
            Currency: "USD"
        }

        account, err := testQueries.CreateAccount(context.Background(), arg)
        require.NoError(t, err)
        require.NotEmpty(t, account)

        require.Equal(t, arg.Owner, account.Owner)
        require.Equal(t, arg.Balance, account.Balance)
        require.Equal(t, arg.Currency, account.Currency)

        require.NotZero(t, account.ID)
        require.NotZero(t, account.CreatedAt)
  }
  ```

  - Most of the data above is hardcoded lets use random utils for this 
    random.go, make util folder and put it inside this
    ```
    package util

    import(
        "math/rand"
        "time"
        "strings"
    )

    const alphabet = "abcdefghijklmnopqrstuvwxyz"

    func init () {
        rand.Seed (time.Now().UnixNano())
    }

    func RandomInt (min, max int64) int64 {
        return min + rand.Int63n(max-min+1)   // between min and max
    }

    func RandomString(n int) string {
        var sb strings.Builder
        k := len(alphabet)

        for i := 0; i < n; i++ {
            c := alphabet[rand.Intn(k)]
            sb.WriteByte(c)
        }

        return sb.String()
    }

    func RandomOwner() string {
        return RandomString(6)
    }

    func RandomMoney() int64 {
        return RandomInt(0, 1000)
    }

    func RandomCurrency() string{
        currencies := []string{"EUR", "USD", "CAD"}
        n := len(currencies)
        return currencies[rand.Intn(n)]
    }
    ```

- Lets go back and make use of the random values
    ```
    import ("github.com/pratyush/simplebank/util")  // import the utils pckg
        arg:= CreateAccountParams{
            Owner: util.RandomOwner(),
            Balance: util.RandomMoney(),
            Currency: util.RandomCurrency()
        }    
    ```

- Add test command to make 
    ```
    test:
        go test -v -cover ./...
    ```

    -v for verbose logging
    -cover for coverage report
    ./... for running all scripts at once


## Unit testing for GetAccount
- In the same file include the below codes

```
func TestCreateAccount(t *testing.T){ 
    createRandomAccount(t)
}
// Since to test get account we need to create an account first so this above function helps in separating business logic from other tests, note the above isnt an unit test its just a function
func TestGetAccount(t *testing.T) {
    account1 := createRandomAccount(t) //Assign
    account2, err := testQueries.GetAccount(context.Background(), account1.ID) //Act, account2 is crated as a result

    require.NoError(t, err) //Assertions
    require.NotEmpty(t, account2)

    require.Equal(t, account1.ID, account2.ID)
    require.Equal(t, account1.Owner, account2.Owner)
    require.Equal(t, account1.Balance, account2.Balance)
    require.Equal(t, account1.Currency, account2.Currency)
    require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)  // makes sure these accs are created within the duration
}
```

## Testing UpdateAccount
```
func TestUpdateAccount(t *testing.T) {
    account1 := createRandomAccount(t)
    
    arg := UpdateAccountParams{
        ID: account1.ID,
        Balance: util.RandomMoney(),
    }

    account2, err := testQueries.UpdateAccount(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, account2)

    require.Equal(t, account1.ID, account2.ID)
    require.Equal(t, account1.Owner, account2.Owner)
    require.Equal(t, arg.Balance, account2.Balance)
    require.Equal(t, account1.Currency, account2.Currency)
    require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)  // makes sure these accs are created within the duration
}
```

## Testing DeleteAccount

```
func TestDeleteAccount(t *testing.T)
{
    account1 := createRandomAccount(t)
    err := testQueries.DeleteAccount(context.Background(), account1.ID)
    require.NoError(t, err)

    account2, err := testQueries.GetAccount(context.Background(), account1.ID)
    require.Error(t, err)
    require.EqualError(t, err, sql.ErrNoRows.Error())
    require.Empty(t, account2)
}
```

## Testing list accounts
```
func TestListAccounts(t *testing T) {
    for i := 0; i < 10 ; i ++ {
        createRandomAccount(t)
        }
    arg := ListAccountsParams{
        Limit: 5,
        Offset: 5,
    }

    accounts, err := testQueries.ListAccounts(context.Background(). arg)
    require.NoError(t, err)
    requireLen(t, accounts, 5)

    for _, account := range accounts{
        require.NotEmpty(t, account)
    }
}
```

## Final Notes
- On running all tests you can see the coverage isnt 100% thats becoz we till have to write tests for entries and transfers, refer to their test scripts to see how to write tests for them.

- We havent used mocking and spies till now, we will cover them in future lectures.