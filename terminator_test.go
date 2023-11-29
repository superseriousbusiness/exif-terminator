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
	"golang.org/x/image/webp"
)

type TerminatorTestSuite struct {
	suite.Suite
}

func (suite *TerminatorTestSuite) TestTerminateKitten() {
	kitten, err := os.Open("./images/kitten.jpg")
	if err != nil {
		panic(err)
	}
	defer kitten.Close()

	stat, err := kitten.Stat()
	if err != nil {
		panic(err)
	}

	originalSize := int(stat.Size())

	out, err := terminator.Terminate(kitten, originalSize, "jpeg")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(originalSize, len(b))

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
	defer sloth.Close()

	stat, err := sloth.Stat()
	if err != nil {
		panic(err)
	}

	originalSize := int(stat.Size())

	out, err := terminator.Terminate(sloth, originalSize, "jpeg")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(originalSize, len(b))

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
	defer comic.Close()

	stat, err := comic.Stat()
	if err != nil {
		panic(err)
	}

	originalSize := int(stat.Size())

	out, err := terminator.Terminate(comic, originalSize, "png")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(originalSize, len(b))

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
	defer turnip.Close()

	stat, err := turnip.Stat()
	if err != nil {
		panic(err)
	}

	originalSize := int(stat.Size())

	out, err := terminator.Terminate(turnip, originalSize, "jpeg")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(originalSize, len(b))

	// should be decodable as a jpeg
	_, err = jpeg.Decode(bytes.NewBuffer(b))
	suite.NoError(err)

	// bytes should be the same as the clean image
	turnipClean, err := os.ReadFile("./images/giant-turnip-world-record-clean.jpg")
	suite.NoError(err)
	suite.EqualValues(turnipClean, b)
}

func (suite *TerminatorTestSuite) TestTerminatePJW() {
	pjw, err := os.Open("./images/pjw.webp")
	if err != nil {
		panic(err)
	}
	defer pjw.Close()

	stat, err := pjw.Stat()
	if err != nil {
		panic(err)
	}

	originalSize := int(stat.Size())

	out, err := terminator.Terminate(pjw, originalSize, "webp")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(originalSize, len(b))

	// should be decodable as a webp
	_, err = webp.Decode(bytes.NewBuffer(b))
	suite.NoError(err)

	// bytes should be the same as the clean image
	pjwClean, err := os.ReadFile("./images/pjw-clean.webp")
	suite.NoError(err)
	suite.EqualValues(pjwClean, b)
}

func (suite *TerminatorTestSuite) TestTerminatePanorama() {
	panorama, err := os.Open("./images/exif-panorama.jpg")
	if err != nil {
		panic(err)
	}
	defer panorama.Close()

	stat, err := panorama.Stat()
	if err != nil {
		panic(err)
	}

	originalSize := int(stat.Size())

	out, err := terminator.Terminate(panorama, originalSize, "jpeg")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(originalSize, len(b))

	// should be decodable as a jpeg
	_, err = jpeg.Decode(bytes.NewBuffer(b))
	suite.NoError(err)

	// bytes should be the same as the clean image
	panoramaClean, err := os.ReadFile("./images/exif-panorama-clean.jpg")
	suite.NoError(err)
	suite.EqualValues(panoramaClean, b)
}

// thanks to kemonine for the test image and the lemon chicken fettuccine recipe ;)
// https://blog.kemonine.info/recipe-lemon-chicken-fettuccine/
func (suite *TerminatorTestSuite) TestTerminateRecipe() {
	recipe, err := os.Open("./images/recipe.jpg")
	if err != nil {
		panic(err)
	}
	defer recipe.Close()

	stat, err := recipe.Stat()
	if err != nil {
		panic(err)
	}

	originalSize := int(stat.Size())

	out, err := terminator.Terminate(recipe, originalSize, "jpeg")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(originalSize, len(b))

	// should be decodable as a jpeg
	_, err = jpeg.Decode(bytes.NewBuffer(b))
	suite.NoError(err)

	// bytes should be the same as the clean image
	recipeClean, err := os.ReadFile("./images/recipe-clean.jpg")
	suite.NoError(err)
	suite.EqualValues(recipeClean, b)
}

func (suite *TerminatorTestSuite) TestTerminateFish() {
	fishDirty, err := os.ReadFile("./images/fish.png")
	if err != nil {
		panic(err)
	}
	originalSize := len(fishDirty)

	// should not be decodable as a png
	_, err = png.Decode(bytes.NewBuffer(fishDirty))
	suite.EqualError(err, "png: invalid format: invalid checksum")

	out, err := terminator.Terminate(bytes.NewBuffer(fishDirty), originalSize, "png")
	suite.NoError(err)

	// we should be able to get some bytes back from the returned reader
	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	// the processed image should have the same size as the initial image
	suite.EqualValues(originalSize, len(b))

	// should be decodable as a png
	_, err = png.Decode(bytes.NewBuffer(b))
	suite.NoError(err)

	// bytes should be the same as the clean image
	fishClean, err := os.ReadFile("./images/fish-clean.png")
	suite.NoError(err)
	suite.EqualValues(fishClean, b)

	// bytes should not be the same as the
	// original, since we fixed some things.
	suite.NotEqual(fishClean, fishDirty)
}

func TestTerminatorTestSuite(t *testing.T) {
	suite.Run(t, &TerminatorTestSuite{})
}
