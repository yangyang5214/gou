package fileutil

//Reference https://github.com/projectdiscovery/utils/blob/main/file/file.go

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"github.com/pkg/errors"
	"github.com/yangyang5214/gou/type"
	"io"
	"net/http"
	"os"
	"strings"
)

var ErrInvalidSeparator = errors.New("invalid separator")

// FileSave is writes the contents of string
func FileSave(filename string, content string) error {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// FileExists checks if the file exists in the provided path
// https://github.com/projectdiscovery/utils/blob/main/file/file.go#L25
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// FolderExists checks if the folder exists
// https://github.com/projectdiscovery/utils/blob/main/file/file.go#L37
func FolderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return info.IsDir()
}

// FileOrFolderExists checks if the file/folder exists
// https://github.com/projectdiscovery/utils/blob/main/file/file.go#L49
func FileOrFolderExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// DownloadFile to specified path
// https://github.com/projectdiscovery/utils/blob/main/file/file.go#L110
func DownloadFile(filepath string, url string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

// CopyFile from source to destination
// https://github.com/projectdiscovery/utils/blob/main/file/file.go#L252
func CopyFile(src, dst string) error {
	if !FileExists(src) {
		return errors.New("source file doesn't exist")
	}
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return dstFile.Sync()
}

// CreateFolders in the list
// https://github.com/projectdiscovery/utils/blob/main/file/file.go#L131
func CreateFolders(paths ...string) error {
	for _, path := range paths {
		if err := CreateFolder(path); err != nil {
			return err
		}
	}

	return nil
}

// CreateFolder path
func CreateFolder(path string) error {
	return os.MkdirAll(path, 0700)
}

// FileWriteLines is a method Write datas to path line by line
func FileWriteLines(path string, lines *[]string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return errors.WithStack(err)
	}

	defer f.Close()
	buf := bufio.NewWriter(f)

	for _, line := range *lines {
		_, err = buf.WriteString(line + "\n")
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return errors.WithStack(buf.Flush())
}

// FileRead 读取文件，遇到 err 返回空
func FileRead(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		return []byte{}
	}
	return content
}

// FileReadLines is read file line by line.
func FileReadLines(path string) *[]string {
	var lines []string
	content := FileRead(path)
	if len(content) == 0 {
		return &lines
	}
	for _, line := range strings.Split(string(content), "\n") {
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			lines = append(lines, line)
		}
	}
	return &lines
}

// FileReadLinesSet is deduplication data by line
func FileReadLinesSet(path string) typeutil.Set {
	s := typeutil.Set{}
	for _, line := range *FileReadLines(path) {
		s[line] = true
	}
	return s
}

// CountLines counts the lines in a file
func CountLines(filename string) (uint, error) {
	return CountLinesWithSeparator([]byte("\n"), filename)
}

// CountLinesWithSeparator of a file
func CountLinesWithSeparator(separator []byte, filename string) (uint, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	if len(separator) == 0 {
		return 0, ErrInvalidSeparator
	}

	return CountLinesWithOptions(file, separator, nil)
}

// CountLinesWithOptions from a reader and custom filter function
func CountLinesWithOptions(reader io.Reader, separator []byte, filter func([]byte) bool) (uint, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)
		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, separator); i >= 0 {
			return i + len(separator), data[0:i], nil
		}
		if atEOF {
			return dataLen, data, nil
		}
		return 0, nil, nil
	})

	var count uint
	for scanner.Scan() {
		if filter == nil || filter(scanner.Bytes()) {
			count++
		}
	}
	return count, scanner.Err()
}
