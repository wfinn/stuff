# semgrep-gener8

This tool can help generate an initial semgrep rule when you want to detect uses of functions without static content in a specific argument.

```
$ echo -e "fmt.Println(0)\nos.Exec(2)" | semgrep-gener8 
rules:
  - id: change_id_here
patterns:
  - pattern-either:
    - fmt.Println(...)
    - os.Exec(...)
  - pattern-not: fmt.Println("...")
  - pattern-not: os.Exec(..., ..., "...")
```
