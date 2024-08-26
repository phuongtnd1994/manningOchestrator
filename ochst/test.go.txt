package main

import (
	"errors"
	"fmt"
)

// CustomError là một loại lỗi tùy chỉnh
type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}

func main() {
	// Tạo một lỗi cơ bản
	baseErr := errors.New("Đây là lỗi cơ bản")

	// Tạo một lỗi tùy chỉnh
	customErr := &CustomError{msg: "Đây là lỗi tùy chỉnh"}

	// Tạo một lỗi gốc từ lỗi tùy chỉnh
	wrappedErr := fmt.Errorf("Đã xảy ra lỗi: %w", customErr)

	// Sử dụng errors.Is để kiểm tra lỗi gốc
	if errors.Is(wrappedErr, customErr) {
		fmt.Println("wrappedErr chứa customErr")
	} else {
		fmt.Println("wrappedErr không chứa customErr")
	}

	// Sử dụng errors.As để kiểm tra loại lỗi
	var targetErr *CustomError
	if errors.As(wrappedErr, &targetErr) {
		fmt.Println("wrappedErr là lỗi CustomError:", targetErr.msg)
	} else {
		fmt.Println("wrappedErr không phải là lỗi CustomError")
	}

	// Kiểm tra với lỗi cơ bản
	if errors.Is(wrappedErr, baseErr) {
		fmt.Println("wrappedErr chứa baseErr")
	} else {
		fmt.Println("wrappedErr không chứa baseErr")
	}
}
