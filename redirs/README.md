# redirs

find (open) redirects in a list of urls

```sh
# install
go install github.com/wfinn/stuff/redirs@latest

# basic usage
cat urls.txt | redirs
```

This tool requests all urls, looks if somethin

## Flags

- -r to set amount amount of go routines
- -c session=abc123 to set a cookie
- -a "Bearer: token" to set the Authorization header
