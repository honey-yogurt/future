package main

import (
	"fmt"
	"reflect"
)

//func main() {
//	syncHarderPoint := NewSyncHarderPoint()
//	handMap := newMessageHandlerDistributor()
//	handMap.registerHandler("syncHarderPoint", "GetId", syncHarderPoint.GetId)
//	handMap.registerHandler("syncHarderPoint", "getNumber", syncHarderPoint.getNumber)
//
//	h := handMap.handler("syncHarderPoint", "GetId")
//	h1 := handMap.handler("syncHarderPoint", "getNumber")
//	fmt.Println(h())
//	fmt.Println(h1())
//}

//func main() {
//	secSyncHarderPoint := sec.NewSecSyncHarderPoint()
//	handMap := newMessageHandlerDistributor()
//	handMap.registerHandler("syncHarderPoint", "GetId", secSyncHarderPoint.GetId)
//	//handMap.registerHandler("syncHarderPoint", "getNumber", secSyncHarderPoint.getNumber)
//	// Cannot use the unexported method 'getNumber' in the current package
//
//	h := handMap.handler("syncHarderPoint", "GetId")
//	//h1 := handMap.handler("syncHarderPoint", "getNumber")
//	fmt.Println(h())
//	//fmt.Println(h1())
//}

//func main() {
//	s := sec.NewSec()
//	fmt.Println(s.S.GetId())
//	handMap := newMessageHandlerDistributor()
//	handMap.registerHandler("syncHarderPoint", "getId", s.S.GetId)
//	h := handMap.handler("syncHarderPoint", "getId")
//	h()
//}

//func main() {
//	point1 := NewSyncHarderPoint()
//	point2 := NewSyncHarderPoint2()
//	fmt.Println(reflect.TypeOf(point1).Kind())
//	fmt.Println(reflect.TypeOf(point2).Kind())
//	handMap := newMessageHandlerDistributor()
//	handMap.registerHandler("syncHarderPoint1", "getId", point1.GetId)
//	handMap.registerHandler("syncHarderPoint1", "getNumber", point1.getNumber)
//	handMap.registerHandler("syncHarderPoint2", "getId", point2.GetId)
//	handMap.registerHandler("syncHarderPoint2", "getNumber", point2.getNumber)
//	h := handMap.handler("syncHarderPoint1", "getId")
//	h1 := handMap.handler("syncHarderPoint1", "getNumber")
//	h2 := handMap.handler("syncHarderPoint2", "getId")
//	h3 := handMap.handler("syncHarderPoint2", "getNumber")
//	fmt.Println(h())
//	fmt.Println(h1())
//	fmt.Println(h2())
//	fmt.Println(h3())
//}

//func main() {
//	syncHarder := NewSyncHarder()
//	handMap := newMessageHandlerDistributor()
//	handMap.registerHandler("syncHarder", "GetId", syncHarder.GetId)
//	handMap.registerHandler("syncHarder", "getNumber", syncHarder.getNumber)
//	h := handMap.handler("syncHarder", "getNumber")
//	fmt.Println(h())
//
//	pp := NewSyncHarderPp()
//	handMap.registerHandler("pp", "GetId", pp.GetId)
//	handMap.registerHandler("pp", "getNumber", pp.getNumber)
//	h1 := handMap.handler("pp", "getNumber")
//	fmt.Println(h1())
//}

//func main() {
//	syncHarder := NewSyncHarder()
//	point := NewSyncHarderPoint()
//	fmt.Println("syncHarder type is", reflect.TypeOf(syncHarder).Kind())
//	fmt.Println("point type is", reflect.TypeOf(point).Kind())
//	sId := syncHarder.GetId
//	fmt.Println("sId type is", reflect.TypeOf(sId).Kind(), "sId value is", sId)
//	pId := point.GetId
//	fmt.Println("pId type is", reflect.TypeOf(pId).Kind(), "pId value is", pId)
//	handMap := newMessageHandlerDistributor()
//	handMap.registerHandler("syncHarder", "GetId", syncHarder.GetId)
//	handMap.registerHandler("point", "GetId", point.GetId)
//	hsId := handMap.handler("syncHarder", "GetId")
//	hpId := handMap.handler("point", "GetId")
//	fmt.Println("hsId type is", reflect.TypeOf(hsId).Kind(), "hsId value is", hsId)
//	fmt.Println("hpId type is", reflect.TypeOf(hpId).Kind(), "hpId value is", hpId)
//
//}

func Add() string {
	return "s"
}

func main() {
	sId := Add
	fmt.Println("sId type is", reflect.TypeOf(sId).Kind(), "sId value is", sId)
	handMap := newMessageHandlerDistributor()
	handMap.registerHandler("syncHarder", "GetId", sId)
	hsId := handMap.handler("syncHarder", "GetId")
	fmt.Println("hsId type is", reflect.TypeOf(hsId).Kind(), "hsId value is", hsId)
}
