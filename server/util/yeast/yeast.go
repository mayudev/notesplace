// yeast: Tiny but linear growing unique id generator.
// https://github.com/unshiftio/yeast
// Go implementation by mayudev

// The MIT License (MIT)

// Copyright (c) 2015 Unshift.io, Arnout Kazemier,  the Contributors.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package yeast

import (
	"math"
	"strings"
	"time"
)

var alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
var length int64 = 64

// Encode returns a string representing the specified number.
func Encode(num int64) string {
	encoded := ""

	for {
		encoded = string(alphabet[num%length]) + encoded
		num = int64(math.Floor(float64(num / length)))

		if num <= 0 {
			break
		}
	}

	return encoded
}

// Decode returns the integer value specified by the given string
func Decode(encoded string) int64 {
	var decoded int64 = 0
	encodedRune := []rune(encoded)

	for i := 0; i < len(encoded); i++ {
		decoded = decoded*length + int64(strings.IndexRune(alphabet, encodedRune[i]))
	}

	return decoded
}

// Generate generates an ID using current timestamp
func Generate() string {
	return Encode(time.Now().UnixMicro())
}
