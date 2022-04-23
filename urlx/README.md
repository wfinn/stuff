# urlx

--get or --set parts of urls

```sh
cat urls.txt | urlx --get path '?' query # get all paths with and queries
cat urls.txt | urlx --get -u host # get unique hostnames
cat urls.txt | urlx --set host=target.tld # change the host of every url
cat urls.txt | urlx --set path=/prefix{{value}} # prefix the path 
```

urlx accepts many variatios for the names of "parts" so you shouldn't have to remember anything.

```sh
# these all do the same
urlx --get hosts < urls.txt
urlx --get hostname < urls.txt
urlx --get domain < urls.txt
# ...
```