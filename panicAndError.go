package golang2

import (
	"log"
	"regexp"
	"runtime"

	"github.com/pkg/errors"
)

//MyPanic 定义一个捕获panic并且输出栈信息的函数
func MyPanic() {
	pan := recover()
	if pan != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		//获取栈的string
		stack := string(buf[:n])
		//屏蔽panic.go和此go文件的stack
		re := regexp.MustCompile(`(?i).*panicAndError\.go.*|.*panic\.go.*|.*MyPanic.*|panic\(.*`)
		stack = re.ReplaceAllString(stack, "")
		re2 := regexp.MustCompile(`(?m)^\s*$`)
		stack = re2.ReplaceAllString(stack, "@@@@@@@@@@@")
		log.Printf("%s\n%s", pan, stack)
	}
}

//MyError 定义一个
func MyError(err error) error {
	return errors.Wrap(err, `[ERROR]`)
}

//MyErrorLog xxx
func MyErrorLog(err error) {
	err = errors.Wrap(err, `[ERROR]`)
	log.Printf("%+v\n", err)
}
