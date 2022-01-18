package terminator_test

import (
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

	err = terminator.Terminate(kitten, int(stat.Size()), "jpeg")
	suite.NoError(err)
}

func (suite *TerminatorTestSuite) TestTerminateSloth() {
	kitten, err := os.Open("./images/sloth.jpg")
	if err != nil {
		panic(err)
	}

	stat, err := kitten.Stat()
	if err != nil {
		panic(err)
	}

	err = terminator.Terminate(kitten, int(stat.Size()), "jpeg")
	suite.NoError(err)
}

func TestTerminatorTestSuite(t *testing.T) {
	suite.Run(t, &TerminatorTestSuite{})
}
