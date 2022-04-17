# redirs

find redirects in a list of urls

```sh
go install github.com/wfinn/stuff/redirs
gau target.tld | grep -E "=/|=http" | uro | redirs
```