# URL Shortner

Simple URL Shortner written in Go

## First Time setup
*
* Create your .env file, please edit it using all your relevant data:
    ```bash
    cp .example.env .env
    ```
* Run `docker compose up` to set up the MYSQL DB
* Create a table using the name you choose in .env using your favorite DB tools
* Run the Go program (`go run main.go`)

## Creating links

Can be done using a POST request to your the /links/create endpoint, e.g.

```bash
curl --location 'localhost:3000/links/create' \
--form 'url="https://www.tweakers.net/"'
```

This will return a JSON response containing the shortend URL