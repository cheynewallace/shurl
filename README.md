# Shurl
Shurl is a GoLang link shortening web service based on a PostgreSQL database.

It provides a basic web interface for adding new custom short URLs and allows you to choose a string between 3 and 8 characters to shorten the long URL by.

(BYO short domain name)

## Usage
Creates web server on port 8080

In development, you can just run `go run *.go` from the directory, then open a browser and hit http://localhost:8080

In production, `go build` the binary first and then run it

### Adding New URLs
Hit `/new` to view the new URL form, enter the long URL and the short URL chars you wish to use, click create.

Example
- Long URL  = http://www.ourlongdomain.com/pages/20160801/something
- Short URL = shur

The route will now match `/shur` and 301 to the Long URL

Hit  and you will be 301'd to the long URL

## Database Configuration
Shurl will connect to your Postgres database using environment variable values.  They are `SHURL_DB_HOST` `SHURL_DB_USER` `SHURL_DB_PASS` and `SHURL_DB_NAME`.
Modify and add the following lines to your `.bashrc` or relevant profile
```
export SHURL_DB_HOST=localhost
export SHURL_DB_USER=shurl
export SHURL_DB_PASS=passwd123
export SHURL_DB_NAME=shurl
```
### Schema
You can create the single `pages` table required by the service by running `pages.sql` found under the sql folder on your database
