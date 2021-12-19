package foo

var PkgVar1 = "pkgVar"

const PkgConst1 = "pkgConst"

type PkgType1 int

type PkgTypeStruct1 struct {
	Field1 string
	Age    int
}

func (s PkgTypeStruct1) MethodOfPkgType1() {
	var localInMethod1 = "MethodOfPkgType:"
	println(localInMethod1, s.Field1, s.Age)
}

func NewPkgType1(i int) *PkgType1 {
	var localInFunc1 = new(PkgType1)
	*localInFunc1 = PkgType1(i)
	return localInFunc1
}
