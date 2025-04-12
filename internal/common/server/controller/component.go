package controller

type componentable interface {
	Name() string
	Init()
	Run()
	Stop()
}
