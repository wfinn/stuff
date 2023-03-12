# semgrep-gener8

This tool can help generate an initial semgrep rule when you want to detect uses of functions without a static string in a specific argument.

```
$ echo -e "fmt.Println(0)\nos.Exec(2)" | semgrep-gener8 
rules:
  - id: change_id_here
    message: change message here
    severity: WARNING
    languages:
      - generic
    patterns:
    - pattern-either:
      - pattern: fmt.Println(...)
      - pattern: os.Exec(...)
    - pattern-not: fmt.Println("...", ...)
    - pattern-not: os.Exec(..., ..., "...", ...)
```
