// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package taglib

import (
	"fmt"
	"os"
	"strings"
	_ "testing"
)

func ExampleDecode() {
	f, err := os.Open("testdata/BeatThee.mp3")
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	tag, err := Decode(f, fi.Size())
	if err != nil {
		panic(err)
	}

	fmt.Println("Title:", strings.TrimSpace(tag.Title()))
	fmt.Println("Artist:", strings.TrimSpace(tag.Artist()))
	fmt.Println("Album:", strings.TrimSpace(tag.Album()))
	fmt.Println("Track:", tag.Track())

	// Output:
	// Title: Beat Thee
	// Artist: Alexander Nakarada
	// Album: Complete Discography (CC BY Attribution 4.0)
	// Track: 189
}
