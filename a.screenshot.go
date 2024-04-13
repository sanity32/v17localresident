package v17localresident

import (
	"image"

	"github.com/vova616/screenshot"
)

type ScreenshotCtl struct{}

type ScreenshotTakeArgs struct {
	Rect image.Rectangle
}

func (scr ScreenshotCtl) Take(arg ScreenshotTakeArgs, resp *image.RGBA) error {
	if rect := arg.Rect; !rect.Empty() {
		r, err := screenshot.CaptureRect(rect)
		*resp = *r
		return err
	}
	r, err := screenshot.CaptureScreen()
	*resp = *r
	return err
}
