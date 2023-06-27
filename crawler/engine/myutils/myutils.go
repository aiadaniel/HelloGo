package myutils

import "os"

func PathExists(path string) (bool, error) {
	_, e := os.Stat(path)
	if e == nil {
		return true, nil
	}
	if os.IsNotExist(e) {
		return false, nil
	}
	return false, e
}

func FileExists(name string) bool {
	//info, _ := os.Stat(name)
	return false
}
