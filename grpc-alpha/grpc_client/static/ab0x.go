// Code generated by fileb0x at "2020-12-17 11:18:45.623956822 -0500 EST m=+0.001095174" from config file "assets.toml" DO NOT EDIT.
// modification hash(82b547c840f4c18d873c93191abd7863.963594ba69d0273dac33db44d8c13043)

package static

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}

// FileServicePem is "service.pem"
var FileServicePem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x95\xb9\x0e\xbb\x5a\xd2\xc4\x73\x9e\xe2\xcb\xd1\x27\xb0\xd9\x83\x1b\x1c\x0e\x07\x83\xd9\x77\x43\x86\xd9\xc1\x18\xb3\x2f\x4f\x3f\xba\xff\x49\xe6\x6a\xa6\x92\x96\x7e\x25\x55\x77\xd2\xaa\xff\xff\x5b\x22\x7a\xa8\xe6\xff\x41\xe4\xfa\xaa\xac\x42\xe0\xa3\x3f\x14\x33\x54\x55\xd6\x5b\x08\xc1\x91\x56\x60\x57\x45\x50\xa9\x81\xe4\x5e\x8e\x70\x97\xed\x65\x96\xe6\x4e\x27\xb8\xb6\x11\xb6\x07\xd9\xcc\x77\x75\xde\x25\x27\x7e\x6a\x43\xa2\xd6\x5b\x66\x02\x07\xe9\x98\xe8\x80\x5d\xf7\x91\x6e\x80\xee\x01\x6e\x01\x12\x6b\x03\x86\xa1\x71\xc0\x0b\x3c\xc5\xca\x0c\x45\x50\x19\xa0\x63\x34\xc3\x45\xbb\xb4\xc7\x52\xe8\x38\x9a\x04\x6a\xc9\xf1\xf6\xca\x8b\x98\x16\xd3\x5b\x50\xc8\x3b\x79\x1a\x12\x3a\x0d\x3f\x3b\xcc\x16\xd1\xc6\xe5\xa6\x7f\x98\xff\x4f\x66\xa0\xff\x0c\x8e\x7d\xf0\x09\x7d\xcc\x70\xe6\x1d\x3a\x7f\x92\x55\x09\x3c\x2d\xaf\x45\xbe\x21\xa2\x3f\xd7\xc0\x6a\xd7\xc2\x47\x78\xe5\xf0\xdf\xdb\xf4\x16\x79\x86\x08\xfe\x78\x60\xdf\x9f\xef\x87\xd0\x62\x71\x74\x0c\x6f\xca\x24\x0d\x55\x85\x6a\x0b\x4c\xb1\xea\xc6\xba\x6b\x1e\xc2\x4e\x8a\xc0\x41\x32\x00\x16\x04\x15\x0f\xfe\xf6\x61\xa5\x41\x50\x21\x30\xdd\x83\xcd\x5e\xf7\xcd\x54\x5e\x3c\xa6\x88\xb4\xbd\x79\x3a\xef\x10\xad\x03\xfb\xa0\x5d\x6a\x72\x71\x04\x73\xcc\xef\x81\x55\x79\xb7\x75\xd3\xda\x56\xc6\xcf\x7a\x0b\x0c\x2a\x13\x1e\x4f\x62\x2e\x9d\x36\xbf\x00\xe9\xeb\x6b\x58\xc8\x4f\xec\x19\x50\x5c\xfa\xc5\xc7\x2c\x70\x1f\xc1\x92\xd9\x35\x18\x9c\xa7\x32\x2e\x30\x0a\x9e\xeb\xd4\x22\xb8\x68\x7d\x29\x17\x95\xb1\xfd\x4c\x98\x72\xf8\x3b\x46\xd1\xb8\x4d\x74\x75\xa8\xdb\xd0\xc8\x2c\x66\x12\xeb\xfd\x1b\x72\x55\x67\x8f\xe7\x46\xc4\xed\xa6\xec\x2d\xaa\x36\x0f\xce\x95\x67\x46\xc5\xa7\x26\x82\xd8\xe0\x03\xf7\x56\x13\xb8\xf5\x2d\x95\xef\x58\xd2\xa4\x1f\x3c\x25\xb1\x50\x9a\x42\x0e\x31\x29\x9e\x53\xe8\x90\x2f\xa6\xd5\x42\x84\x70\x56\x1b\xcb\xf7\x45\xf3\xed\xae\x66\x9e\x7c\xf5\xe7\xce\x5b\x51\x25\x1b\x1d\x8b\x3f\x33\x0a\x26\xe3\xf4\x79\xce\x19\xed\x33\xf4\x95\x97\x1a\xed\x05\x98\xa8\xa0\x85\xbc\x6e\x9c\x58\x6a\x5d\x19\x0c\xb5\x64\x24\xae\xbf\xc9\xb7\xaf\x02\x99\x05\x77\x80\x86\xca\x47\x5f\x76\x3c\x23\xb4\x05\xdf\xa7\xef\x58\x52\xca\x82\xcb\x0d\x74\x34\x42\x99\x54\x3b\x86\xf2\x04\xdc\xc6\xa7\x47\xe4\x73\x12\x09\x1b\x7d\x98\xfd\x36\xa5\xf8\xef\xfb\x99\x7b\xfb\x71\xb9\x74\x42\x73\xd5\x51\x32\x1f\x30\xb4\x36\xae\x6c\x94\x7c\xca\x17\x10\xfb\xd6\xca\xd0\xfa\x61\x71\x0c\xbd\x38\xcb\x42\x25\x82\x0d\x99\xa5\x6b\xcb\x20\x73\xa1\x98\xbb\xd1\x98\xdf\xac\x0b\x06\x81\x90\x0f\x66\x9b\xdc\x78\x01\xfa\xe0\x7a\x87\x66\x7b\xeb\x72\x43\xa5\x64\xf2\xe6\x49\xf4\xd1\x47\xc6\xa4\x02\x58\x86\x3f\x1d\x09\x7b\xb1\xc0\x18\xae\x34\xcb\x16\xd7\x9c\x5b\x55\x5a\xe2\xd7\x2b\x87\x94\xd1\x76\xfa\xed\xbd\xc5\x7e\x8e\x92\x4a\xfe\xf1\x67\x9f\x87\xd0\x60\x53\x3a\xb0\xd4\xfd\x3a\x31\x72\x90\xbf\xf4\xa9\x20\x61\x1a\xa7\xbc\x19\xef\xed\xd3\xa1\x8e\x62\x6e\xf8\x6f\xf0\xca\xde\x38\x1e\x7f\x79\xbc\x0e\xfc\xee\x96\xcc\x31\xf1\x2a\x7d\x96\x09\x78\x87\x7a\x3d\xb2\xf8\x58\x25\x49\x0f\xb1\xd4\x9c\xcf\x15\x37\xfb\x63\xd5\x22\x46\x55\xce\x6c\xdc\xa8\x20\x1f\x7f\xcf\xe7\xb7\x11\x11\x73\x63\x46\xd6\x61\xdc\xe5\x0c\xc1\xae\x3c\x0f\x4d\xe9\x0c\x4a\x69\x91\x0f\x73\x4d\xd4\xcc\x55\xdd\xb0\xd5\x8b\xb6\xb8\xef\x68\xdb\xa7\x6c\xad\x88\xfd\xb9\x52\x21\x49\x9d\xe9\xf4\xe6\x43\x9a\x41\x10\xec\x08\x80\xd4\xd8\x0d\x48\xef\x3a\x88\xa5\xd0\x25\x5d\x11\x06\xbb\x4a\xab\xff\xfc\x85\xfa\x00\xff\x4b\x62\xbd\xbb\xc4\xdf\xd3\x00\xe4\x03\x7a\xe3\xc3\x53\xdf\x94\xe4\x20\x11\xee\x01\x00\xb4\x0a\x31\xe0\x80\x8c\x5e\x3d\xed\x5e\x9d\x13\xda\xea\xb0\x51\xa6\x6b\x9a\x49\xf4\xca\x0f\xdc\x5b\xa4\x2a\x15\x23\x87\x72\x47\x7c\xe6\x1a\x9a\xad\x68\x72\xe9\x68\x9b\x27\x7b\x70\xc3\xeb\xdb\xc8\x0c\x29\xf6\x53\x9e\x9e\xea\x3f\x5c\x52\xd6\x29\x0e\xb1\xb2\x2f\xea\x42\xa8\xb4\x4c\x64\xae\x51\x30\x7b\x7e\x5f\xda\xbb\xcf\xee\x9c\x24\x74\x9d\x7b\x4e\x82\x6a\x58\x87\xb5\x92\xf4\xe0\x17\x8d\xa7\xae\x25\x06\x86\x1b\x6d\xc8\x06\x41\x17\x7c\x9c\xc5\x5c\x03\xad\xba\xa9\xf3\x80\x10\xeb\x99\x23\xbb\x42\x56\x2a\x74\xf2\xe1\xeb\x79\x0f\xf8\x11\xaf\x83\x0f\x93\x7d\x4f\xcd\x17\x07\x57\xfc\x05\x7b\x24\x17\x58\x5d\x96\x76\x23\x09\xac\x70\x58\x15\x70\xcf\xee\x2e\xab\x37\xa5\x78\x5b\xe8\x1d\xcf\x07\x2a\xd4\xc2\xe2\xd6\x3e\x99\xcd\x0f\xec\xa7\x48\xb2\x9e\xc5\x99\x70\x86\xda\xdb\xab\x6a\x4b\x51\xc9\x7f\xb1\x5a\x33\x4b\xc9\xd0\xe0\x98\x8d\xdc\x52\x9c\x5a\x08\x6d\x1f\x56\x5b\xcd\xec\x93\xe4\xbe\xeb\x9c\xe0\x6f\x93\x4a\x7a\xdf\x30\x6a\x86\xc3\xd6\x7f\x22\xba\x15\x8c\x28\x3f\x8e\xc2\x8d\xa7\x26\xff\x61\x59\xc8\x86\x45\x4a\x77\x4d\x65\x69\x80\xde\x47\x4a\x0a\xb6\xa3\xcf\x28\x4d\xd2\xbd\x76\xf1\xf8\x9e\x58\x33\xc8\x4d\x41\x9e\xa6\xe5\x7d\xf4\x1e\xe0\x31\xdc\xd6\x72\x09\xc2\x91\xae\x0b\xcf\x63\x30\x97\x72\x58\xdf\xd1\x46\x44\xef\x72\x43\xab\xf6\x11\x06\x92\xfb\x66\x58\xf8\x16\x07\x05\xbf\x70\x57\x4a\xb7\xa5\x47\x01\x92\x96\xc2\xee\x05\x22\x26\xb3\x58\x0a\x15\x39\x86\x9c\x0c\xe3\x3d\xc5\x7c\xdc\xa8\xbd\x3d\x71\x7c\x74\x8d\x22\x5e\x54\xbf\x0f\x84\xfb\xd8\xf1\x89\xee\xdc\x0f\x35\xdc\xd6\xc8\xfb\x76\x4d\x1b\x19\xa4\xd7\x12\x2c\x1e\xc4\x70\x4b\xd2\x52\x0e\x15\x44\xac\xca\xd6\x7c\xb0\x89\x47\x24\x49\x8f\xd5\x6f\xbf\x02\x62\xfa\x92\xf6\xd0\xbc\xc6\x6d\x37\xe6\xdb\xe7\x95\xf1\x97\xcb\x5a\xf1\xb1\x8e\x87\xf0\x01\x33\x2b\x9a\x1b\x29\x7c\x42\x47\xa8\x13\x1a\xa6\x29\x9f\x10\xa6\x87\xe9\x0c\x31\xb2\x36\xfc\x10\x93\xc0\xbc\xe9\xa0\xfd\x8e\x04\x0e\x34\x7a\xce\xaa\xf6\x1b\x18\x47\xc2\xae\xc5\x40\xa5\xb4\x10\xb0\xe2\xdd\x19\x6a\x2a\xd2\x41\xb2\x6b\x99\x71\x77\xbf\xcb\x2f\x6f\x36\xac\x69\x10\x1f\x71\x30\x27\x8a\x24\x4e\x17\x2e\x62\x05\x4b\x79\x6d\x3f\x99\xba\x97\x19\x9e\x72\xd7\x47\xd8\x63\x05\x59\xf8\x7e\xce\xe3\x0b\xfc\xf5\x17\xf6\xa7\x11\x91\x29\xfd\x77\x4b\xfe\x2b\x00\x00\xff\xff\x3b\xac\xf4\x73\x42\x07\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileServicePem)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "service.pem", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	path = hfs.Prefix + path

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
