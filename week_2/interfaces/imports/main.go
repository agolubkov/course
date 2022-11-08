package main

import "imports/decorator"

func main() {
	var test decorator.ImportedDecorator
	test.CallPrivateMethodsHere()
	test.CallPublicMethodsHere()
	// test.???? спробувати викликати публічні методи iTest
	// test.???? спробувати викликати публічні методи ITest
}
