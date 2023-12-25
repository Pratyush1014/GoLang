# Pouring image in docker
- To list all the running containers (shows a table)
    ```
    docker ps
    ```
- To list all available images
    ```
    docker images
    ```
- Put your first image
    Link: https://hub.docker.com/_/postgres
    ```
    docker pull postgres:16-alpine
    ```
    Use the above command to pour image for postgres sql. This command pulls the latest stable release doesnt pull beta editions.

    General syntax to pull an image in docker 
    docker pull <image>:<tag>
    *Alpine is a light weight edition of the image
    
- Running the image container(container is a running instance of image, a single image can be used to run multiple container)
    ```
    docker run --name<container_name> -e <environment_variable> -p <host_ports:container_ports> -d <image>:<tag>
    ```
    '-p' is used to create bridge between container and host since they run on different ports. This facilitates mapping between host and container.

    ```
    docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine
    ```
    use docker ps to see if the container is running or not.

- Now the postgres server is ready lets try to connect and access the console
    ```
    docker exec -it <container_name_or_id> <command> [args]
    ```
    This allows to run one specific command inside a running container
    ```
    docker exec -it postgres16 psql -U root
    ```
    psql to access console and -U root to connect to root.

    Some psql commands
    - select now(); -> to show current time
    - \q -> to exit from console

    ```
    docker logs <container_name_or_id>
    ```
    this is used to log container actions i.e. the commands executed on console.
    e.g.
    ```
    docker logs postgres16
    ```

#   TablePlus 
- provides an easy way to interact with the databases.
- install it and add a new PostGres engine connection.
- configuration
        Host: 127.0.0.1
        port: 5432
        User: root
        Password: secret
        database: root

        Then hit test and make sure connection successful pops up.

- Paste the schema query generated for postgres sql and run and refresh you shall see all the tables created.
        path: 1. Creating DB Schema\Simple Bank.sql