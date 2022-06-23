/*
   exif-terminator
   Copyright (C) 2022 SuperSeriousBusiness admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"io"
	"os"

	terminator "github.com/superseriousbusiness/exif-terminator"
)

func main() {
	// open a file
	sloth, err := os.Open("./images/sloth.jpg")
	if err != nil {
		panic(err)
	}
	defer sloth.Close()

	// get the length of the file
	stat, err := sloth.Stat()
	if err != nil {
		panic(err)
	}

	// terminate!
	out, err := terminator.Terminate(sloth, int(stat.Size()), "jpeg")
	if err != nil {
		panic(err)
	}

	// read the bytes from the reader
	b, err := io.ReadAll(out)
	if err != nil {
		panic(err)
	}

	// save the file somewhere
	if err := os.WriteFile("./images/sloth-clean.jpg", b, 0666); err != nil {
		panic(err)
	}
}
