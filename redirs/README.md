# redirs

find (potentially open) redirects in a list of urls

- ignores "lame" redirects (http -> https, home, error page)
- can filter for URLs that look like a typical redirect
- marks especially interesting redirects (experimental, filter with `--interesting`)

`redirs` sends GET requests to the provided URLs, but nothing else.  
Consider this preprocessing for actual open redirect hunting.

```sh
# install
go install github.com/wfinn/stuff/redirs@latest

# basic usage
cat urls.txt | redirs
```

## Flags

- -r to set amount amount of go routines
- -c session=abc123 to set a cookie
- -a "Bearer: token" to set the Authorization header
- --typical to ignore URLs that do not look like typical redirects
- --interesting to only output URLs marked as interesting
