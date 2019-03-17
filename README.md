[![Build Status](https://travis-ci.org/corona10/goimagehash.svg?branch=master)](https://travis-ci.org/corona10/goimagehash)
[![GoDoc](https://godoc.org/github.com/corona10/goimagehash?status.svg)](https://godoc.org/github.com/corona10/goimagehash)
[![Go Report Card](https://goreportcard.com/badge/github.com/corona10/goimagehash)](https://goreportcard.com/report/github.com/corona10/goimagehash)
[![Coverage Status](https://coveralls.io/repos/github/corona10/goimagehash/badge.svg)](https://coveralls.io/github/corona10/goimagehash)

# goimagehash
> Inspired by [imagehash](https://github.com/JohannesBuchner/imagehash)

A image hashing library written in Go. ImageHash supports:
* [Average hashing](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)
* [Difference hashing](http://www.hackerfactor.com/blog/index.php?/archives/529-Kind-of-Like-That.html)
* [Perception hashing](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)
* [Wavelet hashing](https://fullstackml.com/wavelet-image-hash-in-python-3504fdd282b5) [TODO]

## Installation
```
go get github.com/corona10/goimagehash
```
## Special thanks to
* [Haeun Kim](https://github.com/haeungun/)

## Usage

``` Go
func main() {
        file1, _ := os.Open("sample1.jpg")
        file2, _ := os.Open("sample2.jpg")
        defer file1.Close()
        defer file2.Close()

        img1, _ := jpeg.Decode(file1)
        img2, _ := jpeg.Decode(file2)
        hash1, _ := goimagehash.AverageHash(img1)
        hash2, _ := goimagehash.AverageHash(img2)
        distance, _ := hash1.Distance(hash2)
        fmt.Printf("Distance between images: %v\n", distance)

        hash1, _ = goimagehash.DifferenceHash(img1)
        hash2, _ = goimagehash.DifferenceHash(img2)
        distance, _ = hash1.Distance(hash2)
        fmt.Printf("Distance between images: %v\n", distance)
        hash3, _ = goimagehash.AverageHashExtend(img1, 16)
        hash4, _ = goimagehash.AverageHashExtend(img2, 16)
        distance, _ = hash3.Distance(hash4)
        fmt.Printf("Distance between images: %v\n", distance)
        fmt.Printf("hash3 bit size: %v\n", hash3.Bits())
        fmt.Printf("hash4 bit size: %v\n", hash4.Bits())

        var b bytes.Buffer
        foo := bufio.NewWriter(&b)
        _ = hash4.Dump(foo)
        foo.Flush()
        bar := bufio.NewReader(&b)
        hash5, _ := goimagehash.LoadImageHashExtend(bar)
}
```

## Release Note

### v0.4.0(TBA)
- **IMPORTANT** goimagehash v0.0.4 does not have compatible with the before
version when it comes to the serialization.

### v0.3.0
- Support DifferenceHashExtend.
- Support AverageHashExtend.
- Support PerceptionHashExtend by @TokyoWolFrog.

### v0.2.0
- Perception Hash is updated.
- Fix a critical bug of finding median value.

### v0.1.0
- Support Average hashing
- Support Difference hashing
- Support Perception hashing
- Use bits.OnesCount64 for computing Hamming distance by @dominikh
- Support hex serialization methods to ImageHash by @brunoro
