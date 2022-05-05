# urlx

extract parts from urls

```sh
cat urls.txt | urlx path # get all paths
cat urls.txt | urlx -u host # get unique hostnames
cat urls.txt | urlx proto :// newhost.tld path '?' query '#' hash # change the host of every url
```

urlx accepts many variatios for the names of "parts" so you shouldn't have to remember anything.