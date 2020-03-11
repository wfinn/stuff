# shmd: shell to markdown
Use markdown in comments of shell scripts.

Then you can convert a script to markdown with shmd.

This basically justs puts everything but comments inside ```.
You can use # before every line or : ' to begin and ' to end a comment block, they won't be visible.

Might be useful for things like bug reports.

```sh
shmd < script.sh > script.md
```

Ideas:
- Use <<something for special blocks? or <<htmltag?
- Rewrite in shell
