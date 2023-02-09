# Go Multi Progress Bar
![maintained](https://img.shields.io/badge/maintained-yes-brightgreen.svg)
![Programming Language](https://img.shields.io/badge/language-Go-orange.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/IceflowRE/go-multiprogressbar/blob/master/LICENSE.md)

[![Go report card](https://goreportcard.com/badge/github.com/IceflowRE/go-multiprogressbar)](https://goreportcard.com/report/github.com/IceflowRE/go-multiprogressbar)
[![Go Reference](https://pkg.go.dev/badge/github.com/IceflowRE/go-multiprogressbar.svg)](https://pkg.go.dev/github.com/IceflowRE/go-multiprogressbar)

---

Go multi progress bar wraps around [schollz/progressbar](https://github.com/schollz/progressbar).

The library might work for simple cases only and is not tested extensively.

The API might introduce breaking changes from commit to commit for now.

## Why another multi progress bar?

Iam aware of the working and tested multi progress bars out there. Use them if you want a battle-tested library.

I am using `schollz/progressbar` heavily, because it is simple and covers all my cases, except multiple progress bars.

## Installation

go-multiprogressbar is compatible with modern Go releases and modules enabled.

```shell
go get github.com/IceflowRE/go-multiprogressbar
```

will resolve and add the package to the current development module, along with its dependencies.

Or just import it and run `go get` afterwards

```go
import "github.com/IceflowRE/go-multiprogressbar"
```

## Usage

### Import

```go
import "github.com/IceflowRE/go-multiprogressbar"
```

### Basic usage

```go
// create a new multi progressbar with default output to os.Stdout
mpb := multiprogressbar.New()
// add progressbars
for _, pBar := range []*progressbar.ProgressBar{
    progressbar.New(150),
    progressbar.New(200),
    progressbar.New(250),
} {
    mpb.Add(pBar)
}

for val := 0; val < 300; val++ {
    time.Sleep(10 * time.Millisecond)
    barId := val % 3
    mpb.Get(barId).Add(1)
}
```

## Contributing

Every contribution and talk about the structure and organization of the project are always welcome.

Commit titles must follow this pattern `[<category>] <Good Description>` e.g. `[gen] Fix file permissions`.

### Testing

Nothing to see here.

## Thanks

A huge thanks to [@schollz](https://github.com/schollz) for this simple and still highly configurable progressbar library!

## MIT License

Copyright 2023-present Iceflower S (iceflower@gmx.de)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
