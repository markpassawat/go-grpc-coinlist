package util

import "fmt"

func Must(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
