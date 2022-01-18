package terminator

import (
	"bufio"
	"fmt"
	"io"

	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
)

func Terminate(in io.Reader, fileSize int, mediaType string) (io.Reader, error) {
	switch mediaType {
	case "jpeg", "jpg":
		return terminateJpeg(in, fileSize, mediaType)
	default:
		return nil, fmt.Errorf("mediaType %s cannot be processed", mediaType)
	}
}

func terminateJpeg(in io.Reader, fileSize int, mediaType string) (io.Reader, error) {
	pipeReader, pipeWriter := io.Pipe()

	v := &jpegVisitor{
		writer: pipeWriter,
	}

	js := jpegstructure.NewJpegSplitter(v)
	v.js = js

	s := bufio.NewScanner(in)
	s.Buffer([]byte{}, fileSize)
	s.Split(js.Split)

	go func() {
		for s.Scan() {
		}
		pipeWriter.Close()
	}()

	return pipeReader, nil
}
