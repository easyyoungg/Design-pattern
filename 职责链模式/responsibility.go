package main

type Handle interface {
	SetSeccessor(h Handle)
	HandleRequest(request int)
}

type ConcreteHandleA struct {
	Handle
}

type ConcreteHandleB struct {
	Handle
}

type ConcreteHandleC struct {
	Handle
}

func (ca *ConcreteHandleA) SetSeccessor(h Handle) {
	ca.Handle = h
}

func (cb *ConcreteHandleB) SetSeccessor(h Handle) {
	cb.Handle = h
}

func (cc *ConcreteHandleC) SetSeccessor(h Handle) {
	cc.Handle = h
}

func (ca *ConcreteHandleA) HandleRequest(request int) {
	if request < 10 {
		println("use HandleA.")
	} else {
		ca.Handle.HandleRequest(request)
	}
}

func (cb *ConcreteHandleB) HandleRequest(request int) {
	if request >= 10 && request < 20 {
		println("use HandleB.")
	} else {
		cb.Handle.HandleRequest(request)
	}
}

func (cc *ConcreteHandleC) HandleRequest(request int) {
	if request > 20 {
		println("use HandleC.")
	} else {
		cc.Handle.HandleRequest(request)
	}
}

func main() {
	ca := ConcreteHandleA{}
	cb := ConcreteHandleB{}
	cc := ConcreteHandleC{}

	ca.SetSeccessor(&cb)
	cb.SetSeccessor(&cc)

	list := []int{1, 11, 21}

	for _, num := range list {
		ca.HandleRequest(num)
	}
}
