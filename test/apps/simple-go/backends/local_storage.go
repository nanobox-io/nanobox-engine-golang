package backends

import (
	"io"
	"os"
	"io/ioutil"
	"regexp"
	"strings"
)

type LocalStorage struct {
	path string
}

func NewLocalStorage(path string) *LocalStorage {
	return &LocalStorage{
		path: path,
	}
}

func (self *LocalStorage) WriteCloser(id string) (io.WriteCloser, error) {
	err := os.MkdirAll(self.dirPath(id), 0777)
	if err != nil {
		return nil, err
	}
	return os.Create(self.fullPath(id))
}

func (self *LocalStorage) ReadCloser(id string) (io.ReadCloser, error) {
	// get the path and filename
	return os.Open(self.fullPath(id))
}

func (self *LocalStorage) Move(from, to string) error {

	return os.Rename(self.fullPath(from), self.fullPath(to))
}

func (self *LocalStorage) Delete(id string) error {
	// get the path and filename
	err := os.Remove(self.fullPath(id))
	if err == nil || strings.Contains(err.Error(), "no such file") {
		go func() {
			dir := self.dirPath(id)
			re := regexp.MustCompile(`/\w*$`)

			for i := 1; i < len(strings.Split(id, "-")); i++ {
				if emptyFolder(dir) {
					os.Remove(dir)
				}
				dir = re.ReplaceAllString(dir, "")
			}
		}()
	}
	return err
}

func (self *LocalStorage) FileExists(id string) bool {
	f, err := os.Open(self.fullPath(id))
	if err != nil {
		return false
	}
	defer f.Close()
	_, err = f.Stat()
	if err != nil {
		return false
	}
	return true
}

func (self *LocalStorage) dirPath(id string) string {
	re := regexp.MustCompile(`/\w*$`)
	return re.ReplaceAllString(self.fullPath(id), "")
}

func (self *LocalStorage) fullPath(id string) string {
	re := regexp.MustCompile("-")
	return self.path + "/" + re.ReplaceAllString(id, "/")
}

func emptyFolder(dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil || len(files) > 0 {
		return false
	}
	return true
}
