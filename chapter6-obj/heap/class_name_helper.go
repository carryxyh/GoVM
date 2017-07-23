package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

// [XXX => [XXX
// int  => I
// XXX  => LXXX;
func toDescriptor(className string) string {
	if className[0] == '[' {
		// array
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		// primitive
		return d
	}
	// object
	return "L" + className + ";"
}