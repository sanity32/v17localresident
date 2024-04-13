package v17localresident

import "github.com/go-vgo/robotgo"

type MouseCtl struct{}

type MouseMoveArgs struct {
	X      int
	Y      int
	Smooth bool
}

func (m MouseCtl) Move(args MouseMoveArgs, resp *bool) error {
	switch args.Smooth {
	case true:
		robotgo.MoveSmooth(args.X, args.Y)
	case false:
		robotgo.Move(args.X, args.Y)
	}
	return nil
}

type MouseClickArgs struct {
	Button string
	Double bool
}

func (m MouseCtl) Click(args MouseClickArgs, resp *bool) error {
	robotgo.Click(args.Button, args.Double)
	if resp != nil {
		*resp = true
	}
	return nil
}

func (m MouseCtl) Location(args struct{}, resp *[2]int) error {
	x, y := robotgo.Location()
	*resp = [2]int{x, y}
	return nil
}
