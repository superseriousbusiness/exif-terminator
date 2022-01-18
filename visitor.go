package terminator

import (
	"fmt"
	"io"

	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
)

type jpegVisitor struct {
	js     *jpegstructure.JpegSplitter
	writer io.Writer
}

func (v *jpegVisitor) HandleSegment(markerId byte, markerName string, counter int, lastIsScanData bool) error {
	segmentList := v.js.Segments()
	segments := segmentList.Segments()
	lastSegment := segments[len(segments)-1]

	fmt.Println(markerId, markerName, counter, lastIsScanData, lastSegment.IsExif())
	if lastSegment.IsExif() {
		zeroed := make([]byte, len(lastSegment.Data))
		fmt.Printf("writing %d bytes over exif data of length %d\n", len(zeroed), len(lastSegment.Data))
		n, err := v.writer.Write(zeroed)
		fmt.Printf("wrote %d\n", n)
		return err
	}
	

	n, err := v.writer.Write(lastSegment.Data)
	fmt.Printf("wrote %d\n", n)
	return err
}
