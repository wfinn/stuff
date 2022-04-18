# redirs

find redirects in a list of urls

```sh
# install
go install github.com/wfinn/stuff/redirs@latest

# basic usage
cat urls.txt | redirs

# get some urls to scan
gau target.tld | urlame > urls.txt
```

## Flags

- -r to set amount amount of go routines
- -c session=abc123 to set a cookie
- -a "Bearer: token" to se tteh Authorization header