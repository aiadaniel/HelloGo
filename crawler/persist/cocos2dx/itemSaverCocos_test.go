package cocos2dx

import (
	"HelloGo/crawler/parsers/cocos2dx/model"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestItemSaverCocos(t *testing.T) {
	h := ReadTemplate("py_XXX.h")
	log.Printf("%s", h)
	log.Printf("\n")
	log.Printf("**********************************************")
	log.Printf("\n")
	cpp := ReadTemplate("py_XXX.cpptemplate")
	log.Printf("%s", cpp)
}

func TestCreateFile(t *testing.T) {
	h := ReadTemplate("py_XXX.h")

	prefix, _ := filepath.Abs(filepath.Dir("py_XXX.h"))
	var seperator = string(os.PathSeparator)
	dirname := prefix + seperator + "out" + seperator + "2D" + seperator
	filename := fmt.Sprintf("py_%s", "Node.h")
	CreateFile(dirname, dirname+filename, h)
}

//const includeFile = "py_XXX.h"
//const cppFile = "py_XXX.cpptemplate"
//const seperator = string(os.PathSeparator)
func TestGenerateFile(t *testing.T) {
	var classes = make(map[string][]model.Memitem)
	k := "2D Nodes,Animation"

	item := model.Memitem{}
	item.ClassName = "Animation"
	item.PackageName = "2D Nodes"
	item.MemitemLeft = "SpriteFrame *"
	item.MemitemCenter = "getSpriteFrame"
	item.MemitemRight = "() const"
	classes[k] = append(classes[k], item)
	item1 := model.Memitem{}
	item1.ClassName = "Animation"
	item1.PackageName = "2D Nodes"
	item1.MemitemLeft = "const ValueMap &"
	item1.MemitemCenter = "getUserInfo"
	item1.MemitemRight = "() const"
	classes[k] = append(classes[k], item1)
	item2 := model.Memitem{}
	item2.ClassName = "Animation"
	item2.PackageName = "2D Nodes"
	item2.MemitemLeft = "void"
	item2.MemitemCenter = "setUserInfo"
	item2.MemitemRight = "const ValueMap &userInfo"
	classes[k] = append(classes[k], item2)

	includeprefix, _ := filepath.Abs(filepath.Dir(includeFile))
	s, _ := filepath.Abs(includeFile)
	ss, _ := filepath.Abs(cppFile)
	h := ReadTemplate(s)
	cpp := ReadTemplate(ss)

	GenerateFile(classes, includeprefix, h, cpp)

}

func TestGetCurrentPath(t *testing.T) {
	log.Printf("%s", filepath.Dir(includeFile)) //实测返回 .即当前目录
	pp, _ := filepath.Abs(filepath.Dir(includeFile))
	log.Printf("%s", pp)
	dir, _ := os.Getwd()
	log.Printf("%s", dir)
	s, e := GetCurrentPath()
	if e != nil {
		panic("eeee")
	}
	log.Printf("%s", s)
}

func Test96(t *testing.T) {
	log.Printf("%s", string(96))
}
