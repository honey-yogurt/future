# 编码风格原则

这里列举了几条有关思考如何编写可读 Go 代码的总体原则，按重要性排序

+ 清晰：代码的目的和原理对读者来说是清晰的。
+ 简单：代码以最简单的方式完成其目标。
+ 简明：代码具有较高的信噪比。
+ 可维护性：编写的代码可以很容易维护。

## 清晰（Clarity）

可读性的核心目标是生产对读者清晰的代码。

清晰主要是通过有效的命名、有用的注释和有效的代码组织来实现的。

清晰与否要从代码的读者角度来看，而不是从代码的作者的角度来看。代码的易读性比易写性更重要。代码的清晰性有两个不同的方面：

+ 代码实际上在做什么？
+ 为什么代码会执行这样的操作？

### 代码实际上在做什么？

Go的设计使得代码的功能相对直观易懂。在存在不确定性或读者需要先前知识才能理解代码的情况下，**值得花时间让代码的目的更清晰明了**，以便未来读者能够更好地理解。例如，可以采取以下措施：

+ 使用更具描述性的变量名
+ 添加额外的注释
+ 用空白和注释来分隔代码
+ 将代码重构为独立的函数/方法，使其更加模块化。

### 为什么代码会执行这样的操作？

代码的理念通常可以通过变量、函数、方法或包的命名来充分传达。如果无法做到这一点，添加注释就显得非常重要了。当代码中存在读者可能不熟悉的细微差别时，“为什么？”尤其重要，比如：

+ 语言中的微妙之处，例如闭包将捕获一个循环变量，但是闭包与其相距多行。
+ 一个业务逻辑的微妙之处，例如需要区分实际用户和冒充用户的访问控制检查。

一个API可能需要小心使用。例如，一段代码可能因为性能原因而复杂且难以理解，或者一系列复杂的数学运算可能以意想不到的方式使用类型转换。在这些情况和许多其他情况下，重要的是附带的评论和文档解释这些方面，以便未来维护人员不会犯错误，并且读者可以理解代码而无需逆向工程。

还要注意，有些提供清晰度的尝试（如添加额外注释）实际上可能通过增加混乱、重申代码已经表达的内容、与代码相矛盾或增加维护负担来模糊了代码的目标。**让代码自己说话**（例如通过使符号名称本身具有自描述性）而不是添加冗余注释通常更好。注释往往更适合解释为什么要做某事，而不是说明代码正在做什么。

Go标准库中包含了许多践行这一原则的实例，其中包括：

+ 在[sort包](https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/sort/sort.go)中的维护者注释；
+ 同一个包中有[好的可运行示例](https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/sort/example_search_test.go)，这对用户（它们会显示在[godoc](https://pkg.go.dev/sort#pkg-examples)中）和[维护者](https://google.github.io/styleguide/go/decisions#examples)（作为测试的一部分运行）都是有益的。
+ [strings.Cut](https://pkg.go.dev/strings#Cut)只有四行代码，但它们提高了[调用点的清晰度和正确性](https://github.com/golang/go/issues/46336)。

## 简单（Simplicity）

你的 Go 代码应该对使用、阅读和维护它的人来说简单明了。

Go 代码应该以最简单的方式编写，既要实现预期的功能，又要保证良好的性能。简洁的代码是指：

+ 阅读起来很简单
+ 不假设你已经知道它在做什么
+ 不假设你能记住所有前面的代码
+ 没有不必要的抽象层次
+ 没有将注意力引向平凡事物的命名方式
+ 清晰地传递值和决策给读者
+ 有解释为什么而非如何执行代码的注释，以避免未来出现差错
+ 具备独立自成体系的文档说明
+ 拥有有用的错误信息和有效的测试失败结果

在代码简洁性和API使用简便性之间可能存在权衡。例如，让代码更复杂以使API的最终用户能够更容易地正确调用该API可能是值得的。相反，也有可能将一些额外的工作留给API的最终用户，以保持代码简单易懂。

当代码需要复杂性时，应该有意识地增加复杂性。这通常是必要的，如果需要额外的性能或者某个特定库或服务有多个不同客户端时。复杂性可能是合理的，但应该配备相关文档，以便客户端和未来维护人员能够理解和处理这种复杂性。此外还应提供测试和示例来演示其正确使用方式，尤其是如果存在“简单”和“复杂”两种使用方式。

这个原则并不意味着不能在Go中编写复杂代码或者禁止Go代码变得复杂。我们致力于创建一个避免不必要复杂度的代码库，**在出现复杂度时表明涉及到了需要仔细理解和维护的问题。理想情况下，应该附带注释来解释原因，并指出所需注意事项**。当对代码进行优化时经常会出现这种情况；为了达到优化目标往往需要采用更复杂的方法，比如预分配缓冲区并在整个goroutine生命周期内重复使用它。当维护者看到这一点时，应该意识到涉及到的代码是性能关键型的，并且这将影响未来变更时所需注意的事项。然而，如果不必要地使用了这种复杂性，则会给未来需要阅读或修改代码的人带来负担。

如果代码在目标明确为简单时却变得非常复杂，通常表明有必要重新审视实现方式，看是否有更简单的方法来完成同样的任务。

### 最小机制（Least mechanism）

在表达相同意思的几种方式中，更倾向于使用最标准的工具。尽管存在复杂的机制，但不应无故采用。根据需要增加代码复杂性很容易，而发现不必要的复杂性后再去除却困难重重。

+ 在满足需求时，尽量使用核心语言结构（例如通道、切片、映射、循环或结构体）。
+ 如果没有合适的核心语言结构，则可以在标准库中寻找相关工具（如HTTP客户端或模板引擎）。
+ 最后，在引入新依赖项或创建自己的依赖之前，请考虑是否代码库中已经存在足够功能的核心库。

举个例子，如果一段代码需要检查集合的成员资格，一个值元素类型为布尔类型的map（例如map[string]bool）往往就足够了。只有在需要更复杂的操作，而使用map不可能完成或完成起来过于复杂时，才应使用提供类似集合类型和功能特性的库。

## 简明（Concision）

简洁的Go代码具有很高的信噪比。可以轻松辨别相关细节，命名和结构引导读者了解这些细节。

在任何时候，都有很多东西会阻碍主要的细节浮现出来：

+ 重复的代码
+ 不相干的语法
+ [难懂的名字](https://google.github.io/styleguide/go/guide#naming)
+ 不必要的抽象
+ 空白

重复的代码尤其模糊了每个几乎相同部分之间的差异，并需要读者在视觉上比较相似的代码行以找到变化。[表驱动测试](https://github.com/golang/go/wiki/TableDrivenTests)是一个很好的例子，可以将公共代码从每次重复中简洁地提取出来，但选择包含在表格中的部分会影响表格易于理解程度。

在考虑多种代码结构方式时，值得考虑哪种方式最能突显重要细节。

理解和使用常见的代码结构和惯用语也对保持高信噪比至关重要。例如，在错误处理中以下代码块非常常见，读者可以迅速理解该块的目的。

```go
// Good:
if err := doSomething(); err != nil {
    // ...
}
```

如果代码看起来与此非常相似，但微妙地不同，读者可能不会注意到变化。在这种情况下，值得有意地通过添加注释来增强错误检查的信号以引起注意。

```go
// Good:
if err := doSomething(); err == nil { // if NO error
    // ...
}
```

## 可维护性（Maintainability）

代码被编辑的次数比编写的次数多得多。可读性强的代码不仅对于试图理解其工作原理的读者有意义，而且对于需要修改它的程序员也很重要。清晰明了是关键。

可维护的代码：

+ 易于被未来的程序员正确修改
+ 具有良好结构化的API，可以优雅地扩展
+ 明确其所做出的假设，并选择与问题结构相匹配的抽象，而不是代码结构
+ 避免不必要的耦合，并且不包含未使用的功能
+ 拥有全面的测试套件，以确保保持承诺行为和重要逻辑正确无误，并在失败时提供清晰可行动态诊断

当使用抽象概念，如接口和类型时，这些概念本质上会从它们所用的上下文中移除信息。因此，确保它们提供足够的好处非常重要。编辑器和集成开发环境（IDE）可以直接连接到方法定义，并在使用具体类型时显示相应的文档，但对于接口定义则只能进行引用。接口是一个强大的工具，但也有一定代价，因为维护者可能需要了解底层实现的具体细节才能正确地使用该接口，在接口文档或调用点必须进行解释。

可维护性代码还避免将重要细节隐藏在容易被忽视的地方。例如，在以下每行代码中，单个字符是否存在都至关重要以理解该行：

```go
// Bad:
// The use of = instead of := can change this line completely.
if user, err = db.UserByID(userID); err != nil {
    // ...
}
```

```go
// Bad:
// The ! in the middle of this line is very easy to miss.
leap := (year%4 == 0) && (!(year%100 == 0) || (year%400 == 0))
```

