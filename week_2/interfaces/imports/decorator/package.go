package decorator

type iTest interface {
	// додайте ваші публічні та приватні методи
}

type ITest interface {
	// додайте ваші публічні методи
}

type ImportedDecorator struct {
	iTest
	ITest
}

func (id *ImportedDecorator) CallPrivateMethodsHere() {
	// викликати приватні методи iTest
}

func (id *ImportedDecorator) CallPublicMethodsHere() {
	// викликати публічні методи iTest
}
