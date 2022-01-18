package terminator_test

import (
	"fmt"
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

	b, err := io.ReadAll(out)
	suite.NoError(err)
	suite.NotEmpty(b)

	fmt.Println(len(b))

	os.WriteFile("test.jpeg", b, 0666)

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

	b := []byte{}
	_, err = out.Read(b)
	suite.NoError(err)

	suite.NotEmpty(b)
}

func TestTerminatorTestSuite(t *testing.T) {
	suite.Run(t, &TerminatorTestSuite{})
}
