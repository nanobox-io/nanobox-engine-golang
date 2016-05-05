package backendtest

import (
	"io"
	"io/ioutil"
	"strings"
)

type BackendRecorder struct {
	Calls   []string
	Written []byte
}

func (self *BackendRecorder) Write(p []byte) (n int, err error) {
	self.Calls = append(self.Calls, "Write:"+string(p))
	self.Written = append(self.Written, p...)
	return len(p), nil
}
func (self *BackendRecorder) Read(p []byte) (n int, err error) {
	self.Calls = append(self.Calls, "Read")
	return 0, io.EOF
}
func (self *BackendRecorder) Close() error {
	self.Calls = append(self.Calls, "Close")
	return nil
}

func (self *BackendRecorder) WriteCloser(id string) (io.WriteCloser, error) {
	self.Calls = append(self.Calls, "WriteCloser:"+id)

	return self, nil
}

func (self *BackendRecorder) ReadCloser(id string) (io.ReadCloser, error) {
	self.Calls = append(self.Calls, "ReadCloser:"+id)
	// get the path and filename
	return ioutil.NopCloser(strings.NewReader("object data")), nil
}

func (self *BackendRecorder) Move(from, to string) error {
	self.Calls = append(self.Calls, "Move:"+from+"->"+to)
	return nil
}

func (self *BackendRecorder) Delete(id string) error {
	self.Calls = append(self.Calls, "Delete:"+id)
	return nil
}

func NewBackendRecorder() *BackendRecorder {
	return &BackendRecorder{
		Calls:   []string{},
		Written: []byte{},
	}
}
