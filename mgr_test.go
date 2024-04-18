package v17localresident

import (
	"testing"
)

func TestNewMgr(t *testing.T) {
	m := NewMgr(DEFAULT_RPC_SERVER_PORT)
	// defer m.Finalize()
	// if err := m.Connect(); err != nil {
	// 	t.Fatal(err.Error())
	// }
	m.Client().MouseMove(5, 5, true)
}
