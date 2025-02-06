package error

import "fmt"

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("运行时错误: %v", r)
		}
	}()

	if b == 0 {
		panic("除数不能为0")
	}

	result = a / b
	return
}
