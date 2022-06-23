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

package terminator_test

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	terminator "github.com/superseriousbusiness/exif-terminator"
)

type TerminatorTestSuite struct {
	suite.Suite
}

func (suite *TerminatorTestSuite) TestTerminateKitten() {
	kitten, err := os.Open("./images/kitten.jpg")
	if err != nil {
		panic(err)
	}

	stat, err := kitten.Stat()
	if err != nil {
		panic(err)
	}

	out, err := terminator.Terminate(kitten, int(stat.Size()), "jpeg")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(stat.Size(), len(b))

	// should be decodable as a jpeg
	_, err = jpeg.Decode(bytes.NewBuffer(b))
	suite.NoError(err)

	// bytes should be the same as the clean image
	kittenClean, err := os.ReadFile("./images/kitten-clean.jpg")
	suite.NoError(err)
	suite.EqualValues(kittenClean, b)
}

func (suite *TerminatorTestSuite) TestTerminateSloth() {
	sloth, err := os.Open("./images/sloth.jpg")
	if err != nil {
		panic(err)
	}

	stat, err := sloth.Stat()
	if err != nil {
		panic(err)
	}

	out, err := terminator.Terminate(sloth, int(stat.Size()), "jpeg")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(stat.Size(), len(b))

	// should be decodable as a jpeg
	_, err = jpeg.Decode(bytes.NewBuffer(b))
	suite.NoError(err)

	// bytes should be the same as the clean image
	slothClean, err := os.ReadFile("./images/sloth-clean.jpg")
	suite.NoError(err)
	suite.EqualValues(slothClean, b)
}

func (suite *TerminatorTestSuite) TestTerminateComic() {
	comic, err := os.Open("./images/comic.png")
	if err != nil {
		panic(err)
	}

	stat, err := comic.Stat()
	if err != nil {
		panic(err)
	}

	out, err := terminator.Terminate(comic, int(stat.Size()), "png")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(stat.Size(), len(b))

	// should be decodable as a png
	_, err = png.Decode(bytes.NewBuffer(b))
	suite.NoError(err)

	// bytes should be the same as the clean image
	comicClean, err := os.ReadFile("./images/comic-clean.png")
	suite.NoError(err)
	suite.EqualValues(comicClean, b)
}

func (suite *TerminatorTestSuite) TestTerminateTurnip() {
	turnip, err := os.Open("./images/giant-turnip-world-record.jpg")
	if err != nil {
		panic(err)
	}

	stat, err := turnip.Stat()
	if err != nil {
		panic(err)
	}

	out, err := terminator.Terminate(turnip, int(stat.Size()), "jpeg")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(stat.Size(), len(b))

	// should be decodable as a jpeg
	_, err = jpeg.Decode(bytes.NewBuffer(b))
	suite.NoError(err)

	// bytes should be the same as the clean image
	turnipClean, err := os.ReadFile("./images/giant-turnip-world-record-clean.jpg")
	suite.NoError(err)
	suite.EqualValues(turnipClean, b)
}

func (suite *TerminatorTestSuite) TestTerminatePanorama() {
	panorama, err := os.Open("./images/exif-panorama.jpg")
	if err != nil {
		panic(err)
	}

	stat, err := panorama.Stat()
	if err != nil {
		panic(err)
	}

	out, err := terminator.Terminate(panorama, int(stat.Size()), "jpeg")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(stat.Size(), len(b))

	// should be decodable as a jpeg
	_, err = jpeg.Decode(bytes.NewBuffer(b))
	suite.NoError(err)

	// bytes should be the same as the clean image
	panoramaClean, err := os.ReadFile("./images/exif-panorama-clean.jpg")
	suite.NoError(err)
	suite.EqualValues(panoramaClean, b)
}

func TestTerminatorTestSuite(t *testing.T) {
	suite.Run(t, &TerminatorTestSuite{})
}
