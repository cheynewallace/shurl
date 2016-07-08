# Shurl
Shurl is a GoLang link shortening web service based on a PostgreSQL database.

It provides a basic web interface for adding new custom short URLs and allows you to choose a string between 3 and 8 characters to shorten the long URL by.

Eg
`http://www.ourlongdomain.com/pages/20160801/something` to `http://do.ma/shur`

(BYO short domain name)

## Database Configuration
Shurl will connect to your Postgres database using environment variable values.  They are `SHURL_DB_HOST` `SHURL_DB_USER` `SHURL_DB_PASS` and `SHURL_DB_NAME`.
Modify and add the following lines to your `.bashrc` or relevant profile
```
export SHURL_DB_HOST=localhost
export SHURL_DB_USER=shurl
export SHURL_DB_PASS=passwd123
export SHURL_DB_NAME=shurl
```
