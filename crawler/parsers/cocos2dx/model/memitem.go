package model

type Memitem struct {
	PackageName   string
	ClassName     string
	Namespace     string
	MemitemLeft   string // 返回值
	MemitemCenter string // 函数名
	MemitemRight  string // 参数
	Memdesc       string // 函数描述，doc
}
