package v17localresident

import (
	"image"
	"net"
	"net/rpc"
	"time"
)

func NewClient(addr string) *Client {
	return &Client{addr: addr}
}

func ConnectClient(addr string) (*Client, error) {
	r := NewClient(addr)
	return r, r.Connect()
}

type Client struct {
	addr      string
	client    *rpc.Client
	connected bool
}

func (cl *Client) Connect() error {
	conn, err := net.Dial("tcp", cl.addr)
	if err != nil {
		return err
	}
	cl.client = rpc.NewClient(conn)
	cl.connected = true
	return nil
}

func (cl *Client) ConnectN(n int, timeout time.Duration) (err error) {
	for i := 0; i < n; i++ {
		if i != 0 {
			time.Sleep(timeout)
		}
		if err = cl.Connect(); err == nil {
			break
		}
	}
	return err
}

func (cl *Client) init() error {
	if !cl.connected {
		if err := cl.Connect(); err != nil {
			return err
		}
	}
	return nil
}

func (cl *Client) MouseMove(x, y int, smooth bool) error {
	return cl.client.Call("Mouse.Move", MouseMoveArgs{X: x, Y: y, Smooth: smooth}, nil)
}

func (cl *Client) MouseClick(button string, double bool) error {
	return cl.client.Call("Mouse.Click", MouseClickArgs{Button: button, Double: double}, nil)
}

func (cl *Client) MouseLocation() (r [2]int, err error) {
	return r, cl.client.Call("Mouse.Location", struct{}{}, &r)
}

func (cl *Client) KeyTap(key string, args ...any) error {
	return cl.client.Call("Key.Tap", KeyboardActionArgs{key, args}, &struct{}{})
}

func (cl *Client) KeyDown(key string, args ...any) error {
	return cl.client.Call("Key.Down", KeyboardActionArgs{key, args}, &struct{}{})
}

func (cl *Client) KeyUp(key string, args ...any) error {
	return cl.client.Call("Key.Up", KeyboardActionArgs{key, args}, &struct{}{})
}

func (cl *Client) KeyType(text string) error {
	return cl.client.Call("Key.Type", text, &struct{}{})
}

func (cl *Client) Screenshot() (r *image.RGBA, err error) {
	return r, cl.client.Call("Screenshot.Take", ScreenshotTakeArgs{}, &r)
}

func (cl *Client) ScreenshotRect(rect image.Rectangle) (r *image.RGBA, err error) {
	return r, cl.client.Call("Screenshot.Take", ScreenshotTakeArgs{Rect: rect}, &r)
}
