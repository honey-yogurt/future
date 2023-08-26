## 指针接收者，指针对象
两个方法的可访问性不一致，但都是在一个包里面
```go
func main() {
	syncHarderPoint := NewSyncHarderPoint()
	handMap := newMessageHandlerDistributor()
	handMap.registerHandler("syncHarderPoint", "GetId", syncHarderPoint.GetId)
	handMap.registerHandler("syncHarderPoint", "getNumber", syncHarderPoint.getNumber)

	h := handMap.handler("syncHarderPoint", "GetId")
	h1 := handMap.handler("syncHarderPoint", "getNumber")
	fmt.Println(h())
	fmt.Println(h1())
}
```
打印结果：
```text
SyncHarderPoint
77
```

## 指针接收者，指针对象
两个方法的可访问性不一致，但**不在**一个包里面，结构体是可访问
```go
func main() {
	secSyncHarderPoint := sec.NewSecSyncHarderPoint()
	handMap := newMessageHandlerDistributor()
	handMap.registerHandler("syncHarderPoint", "GetId", secSyncHarderPoint.GetId)
	//handMap.registerHandler("syncHarderPoint", "getNumber", secSyncHarderPoint.getNumber)

	h := handMap.handler("syncHarderPoint", "GetId")
	//h1 := handMap.handler("syncHarderPoint", "getNumber")
	fmt.Println(h())
	//fmt.Println(h1())
}
```
无法通过编译，因为 getNumber 不可访问，所以无法注入 **Cannot use the unexported method 'getNumber' in the current package**

## 指针接收者，指针对象
两个方法的可访问性不一致，但**不在**一个包里面，结构体是不可访问
可以把私有结构体包装到一个公有结构体中
```go
type Sec struct {
	S *sSecSyncHarderPoint
}

func NewSec() *Sec {
	return &Sec{S: newsSecSyncHarderPoint()}
}
```

```go
func main() {
	s := sec.NewSec()
	fmt.Println(s.S.GetId())
	handMap := newMessageHandlerDistributor()
	handMap.registerHandler("syncHarderPoint", "getId", s.S.GetId)
	h := handMap.handler("syncHarderPoint", "getId")
	h()
}
```
此时，同样只有公有的方法可以注册，另外其实这种方法和上面不同的是，上面可以通过 .getNumber 进行注册，只不过编译不通过，
通过这种包装方式，甚至无法 s.S.getNumber（编译器是无法.出来的）

猜测：通过公开的结构体，把私有结构体的指针暴露出去了。有了地址，肯定可以访问注册公开方法

## 指针接收者，结构体对象
```go
func main() {
	point1 := NewSyncHarderPoint()
	point2 := NewSyncHarderPoint2()
	fmt.Println(reflect.TypeOf(point1).Kind())
	fmt.Println(reflect.TypeOf(point2).Kind())
	handMap := newMessageHandlerDistributor()
	handMap.registerHandler("syncHarderPoint1", "getId", point1.GetId)
	handMap.registerHandler("syncHarderPoint1", "getNumber", point1.getNumber)
	handMap.registerHandler("syncHarderPoint2", "getId", point2.GetId)
	handMap.registerHandler("syncHarderPoint2", "getNumber", point2.getNumber)
	h := handMap.handler("syncHarderPoint1", "getId")
	h1 := handMap.handler("syncHarderPoint1", "getNumber")
	h2 := handMap.handler("syncHarderPoint2", "getId")
	h3 := handMap.handler("syncHarderPoint2", "getNumber")
	fmt.Println(h())
	fmt.Println(h1())
	fmt.Println(h2())
	fmt.Println(h3())
}
```
打印结果
```text
SyncHarderPoint
77
SyncHarderNoPoint
77
```

由此推测，结构体接收者，结构体对象也是可以的

## 结构体接收者，结构体对象
```go
func main() {
	syncHarder := NewSyncHarder()
	handMap := newMessageHandlerDistributor()
	handMap.registerHandler("syncHarder", "GetId", syncHarder.GetId)
	handMap.registerHandler("syncHarder", "getNumber", syncHarder.getNumber)
	h := handMap.handler("syncHarder", "getNumber")
	fmt.Println(h())
}
```
果然是没有问题的，显然 结构体接收者，指针对象也是没问题的
```go
    pp := NewSyncHarderPp()
	handMap.registerHandler("pp", "GetId", pp.GetId)
	handMap.registerHandler("pp", "getNumber", pp.getNumber)
	h1 := handMap.handler("pp", "getNumber")
	fmt.Println(h1())
```

## 本质
首先肯定和是不是值接收值和指针接收者无关：
+ 值接收者：当你使用值接收者时，Go 会在每次方法调用时复制该值。因此，如果你在方法内部更改接收者的值，这个更改不会影响到原始值。
+ 指针接收者：当你使用指针接收者时，Go 会使用原始值的引用，而不是复制值。因此，如果你在方法内部更改接收者的值，这个更改会影响到原始值。

接收者的类型不同体现在方法内部的行为，而不会体现在我们这个场景。
可以推测，无论是结构体注册还是指针注册，注册进入的都是一个可能是一个指针类似的东西。
可以验证一下:
```go
func main() {
	syncHarder := NewSyncHarder()
	point := NewSyncHarderPoint()
	fmt.Println("syncHarder type is", reflect.TypeOf(syncHarder).Kind())
	fmt.Println("point type is", reflect.TypeOf(point).Kind())
	sId := syncHarder.GetId
	fmt.Println("sId type is", reflect.TypeOf(sId).Kind(), "sId value is", sId)
	pId := point.GetId
	fmt.Println("pId type is", reflect.TypeOf(pId).Kind(), "pId value is", pId)
	handMap := newMessageHandlerDistributor()
	handMap.registerHandler("syncHarder", "GetId", syncHarder.GetId)
	handMap.registerHandler("point", "GetId", point.GetId)
	hsId := handMap.handler("syncHarder", "GetId")
	hpId := handMap.handler("point", "GetId")
	fmt.Println("hsId type is", reflect.TypeOf(hsId).Kind(), "hsId value is", hsId)
	fmt.Println("hpId type is", reflect.TypeOf(hpId).Kind(), "hpId value is", hpId)
}
```
打印结果如下：
```text
syncHarder type is struct
point type is ptr
sId type is func sId value is 0x144d40
pId type is func pId value is 0x144d80
hsId type is func hsId value is 0x144d40
hpId type is func hpId value is 0x144d80
```
可以看出，我们注册的是一个 function 类型，具体值表现是 **这个function的引用（地址）** 。
即，对于某个结构体的实例来说，不仅仅有这个结构体的字段值，我们同样可以通过 实例.方法名（注意不带括号，不是调用） 得到这个结构体实例的 function 值。
**这个 function 值持有 调用者的指针，故可以拿到调用者的信息。**

可以把这个值赋值给其他变量，加上()和输入参数就可以调用这个 function 并得到结果值。

**充分体现了 Go 语言中的函数也是类型， 函数值跟其他普通值一样，函数也可以作为参数传递或作为返回值返回。**
