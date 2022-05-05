# urlame

removes lame urls from a list

```sh
go install github.com/wfinn/stuff/urlame@latest
cat urls.txt | urlame
```

The core idea is to normalize urls, ignoring certain info of it.  
One simple thing is, this only looks at the query parameter names, not the values.

e.g. https://localhost/123?x=ABC and http://localhost/456?x=123 both become https://localhost/%N%?x=%P% internally and urlame then only prints the first one.

It also blacklists certain file extensions like .png and.woff2 and with --block-paths you can save yourself some `grep -v`'s with common paths like /wp-content.

---

Inspired by [uro](https://github.com/s0md3v/uro) (which probably does a much better job at this)
