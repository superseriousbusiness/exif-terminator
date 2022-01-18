package terminator

import (
	"fmt"

	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
)

type jpegVisitor struct {
	js *jpegstructure.JpegSplitter
}

func (v *jpegVisitor) HandleSegment(markerId byte, markerName string, counter int, lastIsScanData bool) error {
	segmentList := v.js.Segments()
	segments := segmentList.Segments()
	lastSegment := segments[counter]
	fmt.Println(markerId, markerName, counter, lastIsScanData, lastSegment.IsExif())
	if lastSegment.IsExif() {
		fmt.Println(lastSegment.Data)
	}
	return nil
}
