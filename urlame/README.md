# urlame

removes lame urls from a list

```sh
go install github.com/wfinn/stuff/urlame@latest
cat urls.txt | urlame
```

The core idea is to normalize urls e.g. https://localhost/123?x=ABC and http://localhost/456?x=123 both become https://localhost/%N%?x=%P% internally and urlame then only prints the first one.

It also blacklists certain paths and file extensions like /static/ and .png.

---

Inspired by [uro](https://github.com/s0md3v/uro) (which probably does a much better job at this)
