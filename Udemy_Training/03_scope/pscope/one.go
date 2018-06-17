package pscope

var x string = "Foobar" // package level scoping, DANGER can be resasigned,
var X string = "this one is globally accessible "

func One() string {
	return X + y
}
