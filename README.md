# URL Shortner

Simple URL Shortner written in Go

Still a work in progress, far from done

## First Time setup
* Create your .env file, please edit it using all your relevant data:
    ```bash
    cp .example.env .env
    ```
* Run `docker compose up` to set up the MYSQL DB
* Create a table using the name you choose in .env using your favorite DB tools
* After making sure that the database exists in your MYSQL, run `go run migrate/migrate.go` to create the neccesairy tables
* Run the Go program (`go run main.go`)

## Creating links

Can be done using a POST request to your the /links/create endpoint, e.g.

```bash
curl --location 'localhost:3000/links/create' \
--form 'url="https://www.tweakers.net/"'
```

This will return a JSON response containing the shortend URL path

## TODOs
* ~~Add authentication for the create route~~ *Basic JWT authentication implementation is now done*
* Add more routes, e.g. `/get-all?limit=20`, `POST /delete {shortLink: 'a1b2c3'}` etc. 
* Figure out some kind of front-end admin panel situation (svelte-kit)
* Add more TODOs
