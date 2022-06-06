package MyID

import (
	"bytes"
	"encoding/binary"
	"os"
	"strings"
)

func IntToBytes(n int64) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, n)
	return buf.Bytes()
}

func BytesToInt(b []byte) int64 {
	var n int64
	binary.Read(bytes.NewReader(b), binary.BigEndian, &n)
	return n
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func Write(path string, content []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	n, _ := f.Seek(0, 2)
	if _, err := f.WriteAt(content, n); err != nil {
		return err
	}
	return nil
}

func SmartWrite(path string, content []byte) error {
	var (
		dst      string
		paths    = strings.Split(path, "/")
		pathsLen = len(paths)
	)
	for k, v := range paths {
		if k != pathsLen-1 {
			dst += v + "/"
		}
	}
	var err error
	if !IsExist(dst) {
		if err = os.MkdirAll(dst, os.ModePerm); err != nil {
			return err
		}
		if err = os.Chmod(dst, os.ModePerm); err != nil {
			return err
		}
	}
	return Write(path, content)
}
