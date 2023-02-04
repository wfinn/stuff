# URLitter

`urlitter` is a dumb project I made with ChatGPT (so their License applies etc).  
The idea was to filter out some URLs from a list which are obviously trash.

## Description

This tool takes a list of URLs from standard input and applies the following filters on each URL:
- URLs whose protocol is not http or https are filtered out
- URLs whose path contain "http://" or "https://" are filtered out
- URLs that contain more special characters than expected characters (where expected characters can be matched with the regular expression `[a-fA-F0-9.-/]*`) are filtered out
- URLs that are longer than a specified maximum number of runes are filtered out

## Options

- `--max-runes` maximum number of runes allowed in a URL (default 1000)  
- `--expected-ratio` minimum ratio of expected characters in a URL (default 0.5)