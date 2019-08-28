package persist

import (
	"HelloGo/crawler/examples/cocos2dx/model"
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
	classes[k] = append(classes[item.ClassName], item)
	item = model.Memitem{}
	item.ClassName = "Animation"
	item.PackageName = "2D Nodes"
	item.MemitemLeft = "const ValueMap &"
	item.MemitemCenter = "getUserInfo"
	item.MemitemRight = "() const"
	classes[k] = append(classes[item.ClassName], item)
	item = model.Memitem{}
	item.ClassName = "Animation"
	item.PackageName = "2D Nodes"
	item.MemitemLeft = "void"
	item.MemitemCenter = "setUserInfo"
	item.MemitemRight = "const ValueMap &userInfo"
	classes[k] = append(classes[item.ClassName], item)

	includeprefix, _ := filepath.Abs(filepath.Dir(includeFile))
	h := ReadTemplate(includeFile)
	cpp := ReadTemplate(cppFile)

	GenerateFile(classes, includeprefix, h, cpp)

}
