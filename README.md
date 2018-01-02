ghselfupdate
=======

[![Build Status](https://travis-ci.org/Songmu/ghselfupdate.png?branch=master)][travis]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/ghselfupdate?status.svg)][godoc]

[travis]: https://travis-ci.org/Songmu/ghselfupdate
[license]: https://github.com/Songmu/ghselfupdate/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/Songmu/ghselfupdate

## Description

easy wrapper for "github.com/rhysd/go-github-selfupdate"

## Synopsis

```golang
import (
    "log"

    "github.com/Songmu/ghselfupdate"
)

err := ghselfupdate.Do(vesion)
if err != nil {
    log.Println("binary update failed:", err)
}
```

## See Also

https://github.com/rhysd/go-github-selfupdate

## Author

[Songmu](https://github.com/Songmu)
