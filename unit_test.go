package filetype

import (
	"fmt"
	"os"
	"testing"
)

func TestSimplyDetectFileTypes(t *testing.T) {
	buf, _ := os.ReadFile("C:\\Users\\zen\\Github\\filetype\\fixtures\\sample.gif")
	// 匹配文件类型
	kind, _ := Match(buf)
	if kind == Unknown {
		fmt.Println("Unknown file type")
		return
	}
	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
}
func TestCheckFileCategories(t *testing.T) {
	buf, _ := os.ReadFile("C:\\Users\\zen\\Github\\filetype\\fixtures\\sample.jpg")

	// 检查是否是图片
	if IsImage(buf) {
		fmt.Println("File is an image")
	} else {
		fmt.Println("Not an image")
	}
}
func TestQuerySupportedTypes(t *testing.T) {
	// 检查是否支持某个扩展名
	if IsSupported("jpg") {
		fmt.Println("Extension supported")
	} else {
		fmt.Println("Extension not supported")
	}

	// 检查是否支持某个 MIME 类型
	if IsMIMESupported("image/jpeg") {
		fmt.Println("MIME type supported")
	} else {
		fmt.Println("MIME type not supported")
	}
}

var fooType = NewType("foo", "foo/foo")

// 定义一个新的匹配器
func fooMatcher(buf []byte) bool {
	return len(buf) > 1 && buf[0] == 0x01 && buf[1] == 0x02
}
func TestAddCustomTypesAndMatchers(t *testing.T) {
	// 注册新的匹配器和类型
	AddMatcher(fooType, fooMatcher)

	// 检查是否支持新的扩展名
	if IsSupported("foo") {
		fmt.Println("New supported type: foo")
	}

	// 检查是否支持新的 MIME 类型
	if IsMIMESupported("foo/foo") {
		fmt.Println("New supported MIME type: foo/foo")
	}

	// 尝试匹配新的类型
	fooFile := []byte{0x01, 0x02}
	kind, _ := Match(fooFile)
	if kind == Unknown {
		fmt.Println("Unknown file type")
	} else {
		fmt.Printf("File type matched: %s\n", kind.Extension)
	}
}
func TestFileHeader(t *testing.T) {
	// Open a file descriptor
	file, _ := os.Open("movie.mp4")

	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)

	if IsImage(head) {
		fmt.Println("File is an image")
	} else {
		fmt.Println("Not an image")
	}
}
