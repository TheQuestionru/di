package di

import ()

func Init(dstPtr interface{}, modules ...ModuleFunc) {
	p := NewPackage(modules...)
	g := p.Build()
	g.Fill(dstPtr)
	g.Fill(dstPtr)
}
