package terminator

import (
	"bufio"
	"fmt"
	"io"

	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
)

func Terminate(r io.Reader, fileSize int, mediaType string) error {
	switch mediaType {
	case "jpeg", "jpg":
		return terminateJpeg(r, fileSize)
	default:
		return fmt.Errorf("mediaType %s cannot be processed", mediaType)
	}

}

func terminateJpeg(r io.Reader, fileSize int) error {
	s := bufio.NewScanner(r)

	// Since each segment can be any size, our buffer must allowed to grow as
	// large as the file.
	buffer := []byte{}
	s.Buffer(buffer, fileSize)

	v := &jpegVisitor{}
	js := jpegstructure.NewJpegSplitter(v)
	v.js = js

	s.Split(js.Split)

	for s.Scan() {
	}

	return nil
}
