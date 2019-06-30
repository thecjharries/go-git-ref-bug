# `go-git` Ref Name Issue

There's currently no ref name validation which can cause some interesting problems.

## Background

`git` has a few naming stipulations. I didn't actually realize how detailed they were until I starting messing around with this earlier today.

* [`git-check-ref-format` docs](https://git-scm.com/docs/git-check-ref-format)
* [The Stack Overflow answer I started with](https://stackoverflow.com/a/12093994)
* [A Regex101 sandbox with the regex in Go](https://regex101.com/r/E2TCqU/3/tests)

## `go-git`

https://github.com/src-d/go-git/blob/master/plumbing/reference.go
https://github.com/src-d/go-git/blob/master/references.go
https://github.com/src-d/go-git/blob/master/config/refspec.go

## Asciicast

[![asciicast](https://asciinema.org/a/254346.svg)](https://asciinema.org/a/254346)
