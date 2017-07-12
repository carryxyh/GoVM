package heap

import (
	"GoVM/chapter2-class/classpath"
	"fmt"
	"GoVM/chapter3-cf/classfile"
)

type ClassLoader struct {
	cp       *classpath.Classpath
	//key 是类的完全限定名称
	classMap map[string]*Class
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:        cp,
		classMap:        make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		//类已经被加载过了
		return class
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s] \n", name, entry)
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	//byte转成class结构体
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := chapter3_cf.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

/**
	除了object，每一个类都有且只有一个超类。需要递归调用LoadClass方法加载他的超类
 */
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//todo
}