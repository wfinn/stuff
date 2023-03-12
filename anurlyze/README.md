# anurlyze

Tool to analyze a list of URLs.

This can quess the types of content in query parameters and path segments.

Output looks like:
```
Url: https://example.org/download?id=123&authtoken=5003deb5d66a25a1f8726204d2cf4098 Types: Hash_MD5,Int
Url: https://example.org/file/123?loggedin=true Types: Bool,Int
Url: https://example.org/upload Types: None
...
```

E.g., when you want to find URLs where parameters or path segments contain JSON do:
```sh
cat urls.txt | anurlyze | grep 'Types: .*JSON'
```

Supported types are:
- Int
- Float
- Date
- Bool
- URL
- JSON
- XML
- Hash_(MD5, SHA-1, SHA-256, MD5, SHA-224, SHA-384, SHA-512, Blake2, Whirlpool) e.g. Hash_SHA-256
- IP (IPv4, IPv6)
- Langcode

Coming soon, maybe:
- Filename (using extensions)
- Phone Number
- Email

Types are sorted, so you can grep for multiples by using things like: `grep 'Types: .*Bool.*Hash_MD5'`.

---

This was written with ChatGPT.
