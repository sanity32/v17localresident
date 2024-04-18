package v17localresident

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	v17localresidentbin "github.com/sanity32/v17localresident/bin"
)

const (
	DEFAULT_REPO_URL_WIN_X64 = "http://github.com/sanity32/lrsrv/releases/download/stable/lrsrv-windows-x64.exe"
	DEFAULT_REPO_URL_DEB_X64 = "http://github.com/sanity32/lrsrv/releases/download/stable/lrsrv-debian-x64"
)

var ErrClientNotInit = errors.New("client is not initialized")

func defFilename() string {
	if runtime.GOOS == "windows" {
		return "lrsrv-windows-x64.exe"
	}
	return "lrsrv-debian-x64"
}

func defArchive() v17localresidentbin.EmbedZipFile {
	if runtime.GOOS == "windows" {
		return v17localresidentbin.LRSRV_WIN_X64_STABLE_ZIPPED
	}
	return v17localresidentbin.LRSRV_DEB_X64_STABLE_ZIPPED
}

func NewMgr(port int) *Mgr {
	return &Mgr{
		Port:               port,
		ExecutableFilename: defFilename(),
		archive:            defArchive(),
	}
}

type Mgr struct {
	Port               int
	ExecutableFilename string
	archive            v17localresidentbin.EmbedZipFile
	cmnd               *exec.Cmd
	client             *Client
}

func (m *Mgr) Finalize() {

	if c := m.cmnd; c != nil {
		if p := c.Process; p != nil {
			p.Kill()
		}
	}
}

func (m *Mgr) Client() *Client {
	if c := m.client; c != nil {
		return c
	}
	panic(ErrClientNotInit)
}

func (m Mgr) addr() string {
	return fmt.Sprintf("localhost:%v", m.Port)
}

func (m Mgr) filepath() string {
	return m.ExecutableFilename
}

func (m Mgr) executableFileExist() bool {
	_, err := os.Stat(m.filepath())
	return err == nil || os.IsExist(err)
}

func (m Mgr) deploy() error {
	return m.archive.ExtractFirst(m.filepath())
}

func (m Mgr) prepareExecutableFile() error {
	if !m.executableFileExist() {
		return m.deploy()
	}
	return nil
}

func (m *Mgr) runExecutableFile() error {
	arg := fmt.Sprintf("-port=%v", m.Port)
	m.cmnd = exec.Command("./"+m.filepath(), arg)
	return m.cmnd.Start()
}

func (m *Mgr) Init() error {
	if err := m.prepareExecutableFile(); err != nil {
		return err
	}
	m.client = NewClient(m.addr())
	if err := m.client.Connect(); err != nil {
		if err := m.runExecutableFile(); err != nil {
			return err
		}
	}
	return m.client.ConnectN(3, time.Millisecond*500)
}
