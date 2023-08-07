package main

import "log"

type component interface {
	draw() string
}

type circle struct {
	c component
}

func (c circle) draw() string {
	return "circle "
}

type square struct {
	c component
}

func (s square) draw() string {
	return "square "
}

type decorate struct {
	c component
}

func (d decorate) draw() string {
	return d.c.draw()
}

type redDecorate struct {
	d decorate
}

type blueDecorate struct {
	d decorate
}

func (r redDecorate) draw() string {
	return r.d.c.draw() + "red"
}

func (b blueDecorate) draw() string {
	return b.d.c.draw() + "blue"
}

func main() {
	//var co component
	var c = circle{}
	var s = square{}

	var rs = redDecorate{d: decorate{s}}
	var bs = blueDecorate{d: decorate{s}}
	var rc = redDecorate{d: decorate{c}}
	var bc = blueDecorate{d: decorate{c}}
	//r.draw(co)
	//b.draw(co)
	a := rs.draw()
	b := rc.draw()
	d := bs.draw()
	e := bc.draw()
	log.Println(a, b, d, e)
}
