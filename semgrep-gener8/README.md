# semgrep-gener8

This tool can help generate an initial semgrep rule when you want to detect uses of functions without a static string in a specific argument.

Let's say you know `$DB.Exec`'s first argument is a valid sink, and `$DB.ExecContext`'s second arg is too,
then you could run `semgrep-gener8` and type something like this:

```
$DB.Exec(0)
$DB.ExecContext(1)
```
and get something like:
```
rules:
  - id: change_id_here
    message: change message here
    severity: WARNING
    languages:
      - generic
    patterns:
    - pattern-either:
      - pattern: $DB.Exec(...)
      - pattern: $DB.ExecContext(...)
    - pattern-not: $DB.Exec("...", ...)
    - pattern-not: $DB.ExecContext(..., "...", ...)
```

The result is a rule that checks for occurances of these functions, where the dangerous argument is not a static string.

The code is crap and will break if you try anything else.

## Flags
- `--lang` to set the language in the rule (default: generic)