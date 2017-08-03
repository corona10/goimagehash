[![Build Status](https://travis-ci.org/corona10/goimagehash.svg?branch=master)](https://travis-ci.org/corona10/goimagehash)
[![GoDoc](https://godoc.org/github.com/corona10/goimagehash?status.svg)](https://godoc.org/github.com/corona10/goimagehash)
[![Go Report Card](https://goreportcard.com/badge/github.com/corona10/goimagehash)](https://goreportcard.com/report/github.com/corona10/goimagehash)

# goimagehash
> Inspired by [imagehash](https://github.com/JohannesBuchner/imagehash)

A image hashing library written in Go. ImageHash supports:
* [Average hashing](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)
* [Difference hashing](http://www.hackerfactor.com/blog/index.php?/archives/529-Kind-of-Like-That.html)
* [Perception hashing](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)
* [Wavelet hashing](https://fullstackml.com/wavelet-image-hash-in-python-3504fdd282b5) [TODO]

**Only support 64bits hash.**

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
}
```
