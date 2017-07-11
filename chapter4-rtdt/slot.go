package chapter4_rtdt

type Slot struct {
	//虚拟机规范要求每个元素至少可以容纳一个int或引用，这里存放int，下面ref存放引用
	num int32
	//引用
	ref *Object
}

type Slots []Slot