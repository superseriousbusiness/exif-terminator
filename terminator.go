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
	// to avoid keeping too much stuff in memory we want to pipe data directly
	pipeReader, pipeWriter := io.Pipe()

	// jpeg visitor is where the spicy hack of streaming the de-exifed data is contained
	v := &jpegVisitor{
		writer: pipeWriter,
	}

	// provide the visitor to the splitter so that it triggers on every section scan
	js := jpegstructure.NewJpegSplitter(v)

	// the visitor also needs to read back the list of segments: for this it needs
	// to know what jpeg splitter it's attached to, so give it a pointer to the splitter
	v.js = js

	// we don't know ahead of time how long segments might be: they could be as large as
	// the file itself, so unfortunately we need to allocate a buffer here that's as large
	// as the file
	s := bufio.NewScanner(in)
	s.Buffer([]byte{}, fileSize)

	// use the jpeg splitters 'split' function, which satisfies the bufio.SplitFunc interface
	s.Split(js.Split)

	// scan asynchronously until there's nothing left to scan, and then close the writer
	// so that the reader on the other side knows that we're done
	//
	// due to the nature of io.Pipe, scanning won't actually execute
	// until the pipeReader starts being read by the function
	go func() {
		for s.Scan() {
		}
		pipeWriter.Close()
	}()

	return pipeReader, nil
}
