package terminator_test

import (
	"bytes"
	"image/jpeg"
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
}

func TestTerminatorTestSuite(t *testing.T) {
	suite.Run(t, &TerminatorTestSuite{})
}
