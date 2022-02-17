# redirex

This tool generates tons of bypasses for open redirects, but can also be used for SSRF etc.

`go install github.com/wfinn/stuff/redirex@latest`

`redirex legitdomain.com attacker.com`

```
.attacker.com
https://legitdomain.com@attacker.com
/%09/attacker.com
https://legitdomain.com(.attacker.com
...
```
