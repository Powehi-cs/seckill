package utils

import (
	"github.com/Powehi-cs/seckill/pkg/errors"
	"io/ioutil"
)

var lock string
var unlock string

func InitLua() {
	script, err := ioutil.ReadFile("../../scripts/get_lock.lua")
	errors.PrintInStdout(err)
	lock = string(script)

	script, err = ioutil.ReadFile("../../scripts/release_lock.lua")
	errors.PrintInStdout(err)
	unlock = string(script)
}

func GetPairLock() (string, string) {
	return lock, unlock
}
