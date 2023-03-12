# getallwords (gaw)

getallwords (gaw) generates a wordlist from a list of urls, like from [gau](https://github.com/lc/gau).

```sh
gau target.tld > urls.txt
gaw < urls.txt > wordlist.txt
```

This does not make any requests, it justs splits up the URL into many parts.
