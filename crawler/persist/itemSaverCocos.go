package persist

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/examples/cocos2dx/model"
	"HelloGo/crawler/myutils"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

/**
 * 爬虫收到的解析包使用模板引擎或者字符串直接处理的方式，生成头文件和.cpp文件
 */
const includeFile = "py_XXX.h"
const cppFile = "py_XXX.cpptemplate"
const seperator = string(os.PathSeparator)

//替换常规方法YYY 底部加一空行
const methodArgs = `//// {LLL} {MMM}{RRR}
PyObject* Py{XXX}_{MMM}(Py{XXX} *self{AAA})
{
    // parse args here

    cocos2d::{XXX} *dr = dynamic_cast<cocos2d::{XXX}*>(self->ob_body);
    // add your code here

    Py_RETURN_NONE;
}

`

//替换模块定义MMM
const moduleStart = `PyMethodDef Py{XXX}_methods[] = {`
const moduleArgs = `	{"{MMM}",(PyCFunction)Py{XXX}_{MMM},METH_VARARGS,""},`
const moduleNoArg = `	{"{MMM}",(PyCFunction)Py{XXX}_{MMM},METH_NOARGS,""},`
const moduleEnd = `	{NULL,NULL,0,NULL}
};`

var classes map[string][]model.Memitem //key 包名+类名  value 函数列表

func ItemSaverCocos(index string) (chan engine.Item, error) {
	includeprefix, _ := filepath.Abs(filepath.Dir(includeFile))
	basepath := includeprefix + seperator + "crawler" + seperator + "persist" + seperator //TODO 优化工作目录问题
	s, _ := filepath.Abs(basepath + includeFile)
	ss, _ := filepath.Abs(basepath + seperator + cppFile)
	h := ReadTemplate(s)
	cpp := ReadTemplate(ss)
	classes = make(map[string][]model.Memitem)
	out := make(chan engine.Item)
	go func() {
		itemCnt := 0
	Loop:
		for {
			select {
			case item := <-out:
				var memFunc = item.Payload.(model.Memitem)
				log.Printf("%d, %s", itemCnt, memFunc)
				k := memFunc.PackageName + "," + memFunc.ClassName
				classes[k] = append(classes[k], memFunc)
				//log.Printf("%s  %s",k,memFunc.MemitemCenter)
				itemCnt++
			case <-time.After(15 * time.Second): //目前先简单n秒超时(TODO 增加爬虫结束通知)
				GenerateFile(classes, includeprefix, h, cpp)
				log.Printf("final get %d", len(classes))
				break Loop
			}

		}
	}()
	return out, nil
}

//这个实测是C盘下AppData内的exe所在目录
func GetCurrentPath() (string, error) {
	path, e := exec.LookPath(os.Args[0])
	if e != nil {
		return "", e
	}
	return filepath.Abs(path)
}

// 生成h/cpp文件  注意：运行时目录有所改变，不同于单元测试目录，在main层级即当前工作目录
func GenerateFile(classes map[string][]model.Memitem, includeprefix string, h string, cpp string) {
	for k, v := range classes {
		var temp = strings.Split(k, ",")
		members := v
		packageName := temp[0]
		if packageName == "2D Nodes" {
			packageName = "2d"
		}
		if packageName == "3D Nodes" {
			packageName = "3d"
		}
		if packageName == "UI Nodes" {
			packageName = "ui"
		}
		packageName = strings.ToLower(packageName)
		//1. 根据包名、类名创建文件
		dirname := includeprefix + seperator + "out" + seperator + strings.TrimSpace(packageName) + seperator
		hfilename := dirname + fmt.Sprintf("py_%s.h", temp[1])
		CreateFile(dirname, hfilename, "")
		cppfilename := dirname + fmt.Sprintf("py_%s.cpp", temp[1])
		CreateFile(dirname, cppfilename, "")

		//2. 模板替换，插入相应方法列表
		var hreplace, cppreplace string

		var buf bytes.Buffer
		var modBuf bytes.Buffer
		modBuf.WriteString(moduleStart)
		modBuf.WriteString("\n")
		for _, memFunc := range members {
			//i: 判断返回值
			//ii: 判断参数个数
			funcName := strings.Trim(memFunc.MemitemCenter, " ")
			paramsTrim := strings.Trim(memFunc.MemitemRight, " ")
			returnTrim := strings.Trim(memFunc.MemitemLeft, " ")
			params := strings.Split(paramsTrim, ",")
			method := strings.Replace(methodArgs, "{MMM}", funcName, -1)
			method = strings.Replace(method, "{LLL}", returnTrim, -1)
			if paramsTrim != "() const" && paramsTrim != "() const override" {
				paramsTrim = "(" + paramsTrim + ")"
			}
			method = strings.Replace(method, "{RRR}", paramsTrim, -1)

			if params == nil || params[0] == "const" || params[0] == "void" || params[0] == "" ||
				params[0] == "() const" || params[0] == "() const override" { //无参
				buf.WriteString(strings.Replace(method, "{AAA}", "", -1))
				modBuf.WriteString(strings.Replace(moduleNoArg, "{MMM}", funcName, -1))
				modBuf.WriteString("\n")
			} else {
				buf.WriteString(strings.Replace(method, "{AAA}", ",PyObject *args", -1))
				modBuf.WriteString(strings.Replace(moduleArgs, "{MMM}", funcName, -1))
				modBuf.WriteString("\n")
			}
		}
		modBuf.WriteString(moduleEnd)
		hreplace = strings.Replace(h, "{XXX}", temp[1], -1)

		cppreplace = strings.Replace(cpp, "{BBB}", packageName+"/CC"+temp[1]+".h", -1)
		cppreplace = strings.Replace(cppreplace, "{YYY}", buf.String(), -1)
		cppreplace = strings.Replace(cppreplace, "{ZZZ}", modBuf.String(), -1)

		cppreplace = strings.Replace(cppreplace, "{XXX}", temp[1], -1)
		//3. 写文件
		CreateFile(dirname, hfilename, hreplace)
		CreateFile(dirname, cppfilename, cppreplace)
	}
}

//读取模板
func ReadTemplate(name string) string {
	//dir, err := filepath.Abs(name) //filepath.Dir
	//if err != nil {
	//	log.Printf("error dir file %s ,%v", dir, err)
	//	return "err"
	//}
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		if os.IsPermission(err) {
			log.Printf("error no permission")
		} else {
			log.Printf("error open file %v", err)
		}
		return "err"
	}

	reader := bufio.NewReader(file)
	//len := reader.Buffered()
	//var bytes []byte
	//_, err = reader.Read(bytes) //目前测试读取不成功
	//if err != nil && err != io.EOF {
	//	log.Printf("error read file %s ,%v",dir,err)
	//}
	var buffer bytes.Buffer
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("error read file %s ,%v", name, err)
		}
		//log.Printf(" buf = %s\n", string(buf))
		buffer.Write(buf)
	}
	return string(buffer.String())
}

//根据包名、类名从模板（.h/.cpp）创建文件
func CreateFile(packageName string, classname string, content string) {

	b, _ := myutils.PathExists(packageName)
	if !b {
		err := os.MkdirAll(packageName, os.ModePerm) //递归创建  Mkdir普通创建
		if err != nil {
			log.Printf("error create dir %v", err)
			return
		}
	}

	file, err := os.Create(classname)
	if err != nil {
		log.Printf("error create file %s,%v", classname, err)
		//return
	}

	if content == "" {
		return
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write([]byte(content))
	if err != nil {
		log.Printf("error write file %v", err)
	}
	err = writer.Flush()
	if err != nil {
		log.Printf("error flush file %v", err)
	}

}
