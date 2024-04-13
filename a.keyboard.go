package v17localresident

import (
	"github.com/go-vgo/robotgo"
)

type KeyCtl struct{}

type KeyboardActionArgs struct {
	Key  string
	Args []any
}

// ("i", "alt", "command")
func (k KeyCtl) Tap(arg KeyboardActionArgs, resp *struct{}) error {
	robotgo.KeyTap(arg.Key, arg.Args...)
	return nil
}

func (k KeyCtl) Down(arg KeyboardActionArgs, resp *struct{}) error {
	robotgo.KeyDown(arg.Key, arg.Args...)
	return nil
}

func (k KeyCtl) Up(arg KeyboardActionArgs, resp *struct{}) error {
	robotgo.KeyUp(arg.Key, arg.Args...)
	return nil
}

func (k KeyCtl) Type(arg string, resp *struct{}) error {
	robotgo.TypeStr(arg)
	return nil
}
