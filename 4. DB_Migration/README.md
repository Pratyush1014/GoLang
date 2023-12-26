# DB Migration

## Steps

- Link: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
  Use the above link for detailed documentation on migrate CLI

- Install migrate CLI, use brew or scoop etc depending upon your operating system
    ```
    $ brew install golang-migrate
    ```
    ```
    migrate -version
    migrate -help
    ```
    Some useful commands

- Create a folder simplebank within that create a subdirectory db\migrate
    ```
    mkdir db\migration
    ```

- Create migration files
    ```
    migrate create -ext sql -dir db\migration -seq init_schema
    ```

    Two files will be generated with 'up' and 'down' suffixes

    Up -> These files are used to migrate the db up, these are mostly used while we initiate the db creation or we want to employ any changes.
    These files are executed sequentially according to the prefixes in their names.

    We will put all the db schema creation scripts in this file i.e. GoLang\1. Creating DB Schema\Simple Bank.sql

    Down -> These are used to revert the changes made to the data base and these are executed sequentially in the reverse order of prefixes in their names.

    We will mostly put db deletion scripts or rollback scripts in this file.

    Voila our migration scripts are ready.

- Some usefull docker commands
    ```
    docker ps // lists all running containers
    docker stop <container_name> // i.e. postgres16 stops a container
    docker ps -a // list all containers irrespective of their run state
    docker start postgres16
    ```

- Accessing the postgres container's shell
    ```
    docker exec -it postgres16 /bin/sh //you can also use /bash or any shell installed
    ```

    This allowes you to execute commands in the container directly from the shell.

    1. Creating a db
    ```
    /# createdb --username=root --owner=root simple_bank
    ```
    2. Accessing the db shell
    ```
    /# psql simple_bank
    ```
    3. Running Scripts
    You can exit from the shell and run scripts from outside
    ```
    simple_bank=# \q
    / # dropdb simple_bank
    / # exit
    ```

    ```
    pratyush@LAPTOP-8I7J17JH:~$ docker exec -it postgres16 createdb --username=root --owner=root simple_bank
    pratyush@LAPTOP-8I7J17JH:~$ docker exec -it postgres16 psql -U root simple_bank

    simple_bank=# \q
    ```

- Create a MakeFile outside the db\migration folder this will be used as a script guide by other users to execute certain commands, refer to the file its pretty self explanatory.
    Tips: to search for commands previously executed 
    ```
    history | grep "docker run"  // grep "some string" used for pattern matching
    ```

- Start fresh, stop the running container and delete it
    ```
    docker stop postgres16
    docker rm postgres16
    ```

    Run 
    ```
    make postgres
    ```
    Voila you started a new container using your makefile.

    Similarly to create and drop db you can use commands in the make file i.e.
    ```
    make createdb
    make dropdb // optional
    ```
 - Running first migration script
    Once the db is created you can look for simple_bank database in table plus you would find the tables to be empty
    Lets run the first migration script
    ```
    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
    ```

    Looks overwhelming lets break it down
    ```
    -path db/migration // path to migration scripts
    -database "postgresql://root:secret@localhost:5432/simple_bank // postgresql is the database followed by db env variables
    sslmode=disable // by default postgres doesnt enable ssl so you've to manually provide ssl configs
    -verbose // for verbose logging
    up // finally making the up migration script run
    ```

    You will find a schema_migration table which keeps track of all migrations run on the system and the dirty column specifies if an sepecific version of migration script failed executing or not.

    Finally add the migrateup and migratedown to your makefile.




    