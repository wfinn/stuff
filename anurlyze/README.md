# anurlyze

Tool to analyze a list of URLs.

This can quess the types of content in query parameters.

Output looks like:
```
Url: https://example.org/download?id=123&authtoken=5003deb5d66a25a1f8726204d2cf4098 Types: Int,Hash_MD5
Url: https://example.org/upload Types: None
...
```

E.g., when you want to find URLs where parameters contain JSON do:
```sh
cat urls.txt | anurlize | grep 'Types: .*JSON'
```

Supported types are:
- Int
- Date
- Bool
- URL
- IP
- JSON
- XML
- Hash_(SHA-1, SHA-256, MD5, SHA-224, SHA-384, SHA-512, Blake2, Whirlpool) e.g. Hash_SHA-256

Coming soon, maybe:
- Filename (using extensions)
- Langcode
- Phone Number
- Email

---

This was written with ChatGPT.
