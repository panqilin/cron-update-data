package common

import "log"

/**
	用于抛出，捕获异常，错误信息
 */

//捕获异常
func recoverName() {
	if r := recover(); r != nil {
		log.Println("capture erros messages - ", r)
	}
}
//抛出异常
func throwError(e error) {
	if e != nil {
		panic(e.Error())
	}
}