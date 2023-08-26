# 变量与类型

## 整数

|       类型       |     长度     |                             补充                             |
| :--------------: | :----------: | :----------------------------------------------------------: |
| (unsigned) char  | 至少一个字节 | ASCII 表中的字母，<br />也可以用来保存 -128 到 127 之间的小整数 |
|  (unsigned) int  | 至少两个字节 |                                                              |
| (unsigned) short | 至少两个字节 |                   保证 short 不会比 int 长                   |
| (unsigned) long  | 至少四个字节 |                   保证 long 不会比 int 短                    |
|                  |              |                                                              |

溢出问题：

```c
#include <stdio.h>

int main(void) {
  unsigned char j = 255;
  j = j + 10;
  printf("%u", j); /* 9 */
  unsigned char i - 0;
  i = i - 1;
  printf("%u", i); // 255  
  char m = 127;
  m = m + 10;
  printf("%u", m); /* 4294967177 */
}
```

+ unsigned：溢出重置为**边界值**，然后继续相加
+ 有符号的话：程序基本上会给你一个很大的值，这个值可能变化

**C 并不会在你超出类型的限制时保护你。**

## 浮点数

浮点类型可以表示的数值范围比整数大得多，还可以表示整数无法表示的分数。

使用浮点数时，我们将数表示成**小数乘以 10 的幂**。

|    类型     |        长度        |                           补充                            |
| :---------: | :----------------: | :-------------------------------------------------------: |
|    float    |     至少四字节     | 最小要求是 float 可以表示范围在 10^-37 到 10^+37 之间的数 |
|   double    | 更大（具体看机器） |                                                           |
| long double | 更大（具体看机器） |                                                           |

在我的机器上面：

```
char size: 1 bytes
int size: 4 bytes
short size: 2 bytes
long size: 8 bytes
float size: 4 bytes
double size: 8 bytes
long double size: 16 bytes
```



# 常量

常量的声明与变量类似，不同之处在于常量声明的前面带有 const 关键字，并且你总是需要给常量指定一个值。

通常情况下将常量声明为大写。

```c
const int AGE = 37;
```

另一种定义常量的方式是使用这种语法:

```c
#define AGE 37
```

在这种情况下，你**不需要添加类型**，也不需要使用等于符号 = ，并且可以省略末尾的分号。

C 编译器将会在编译时从声明的值**推断**出相应的类型。



# 运算符

## 算术运算符

二元操作符需要两个操作数

| 操作符 | 名字 | 示例  |
| :----: | :--: | :---: |
|   =    | 赋值 | a = b |
|   +    |  加  | a + b |
|   -    |  减  | a - b |
|   *    |  乘  | a * b |
|   /    |  除  | a / b |
|   %    | 取模 | a % b |

一元操作符只需要一个

| 运算符 |  名字  |    示例    |
| :----: | :----: | :--------: |
|   +    | 一元加 |     +a     |
|   -    | 一元减 |     -a     |
|   ++   |  自增  | a++ or ++a |
|   --   |  自减  | a-- or --a |

a++ 与 ++a 的区别在于：**a++ 在使用 a 之后才自增它的值，而 ++a 会在使用 a 之前自增它的值。**

## 比较运算符

| 运算符 | 名字     | 示例   |
| :----- | :------- | :----- |
| ==     | 相等     | a == b |
| !=     | 不相等   | a != b |
| >      | 大于     | a > b  |
| <      | 小于     | a < b  |
| >=     | 大于等于 | a >= b |
| <=     | 小于等于 | a <= b |



## 复合赋值运算符

| 运算符 | 名字       | 示例   |
| :----- | :--------- | :----- |
| +=     | 加且赋值   | a += b |
| -=     | 减且赋值   | a -= b |
| *=     | 乘且赋值   | a *= b |
| /=     | 除且赋值   | a /= b |
| %=     | 求模且赋值 | a %= b |



## 三目运算符

**<条件> ? <表达式> : <表达式>**

三目运算符的功能与 if/else 条件语句相同，但是它更短，还可以被**内联进表达式**。



## 运算符优先级

按照顺序，优先级从低到高：

- 赋值运算符 =
- 二元运算符 + 和 -
- 运算符 * 和 /
- 一元运算符 + 和 -

运算符还具有关联规则，除了一元运算符和赋值运算符之外，该规则总是**从左到右**的。

括号的优先级比其它任何运算符都要高。



# 条件语句

```c
int a = 1;

if (a == 2) {
  /* do something */
} else if (a == 1) {
  /* 进行一些操作 */
} else {
  /* 进行另一些操作 */
}
```

**数字 0 总是等于 false。其它的任何东西都是 true，包括负数。**

```c
int a = 1;

switch (a) {
  case 0:
    /* 进行一些操作 */
    break;
  case 1:
    /* 进行另一些操作 */
    break;
  case 2:
    /* 进行另一些操作 */
    break;
  default:
    /* 处理所有其它的情况 */
    break;
```

# 循环

```c
for (int i = 0; i <= 10; i++) {
  /* 反复执行的指令 */
}
```

```c
int i = 0;

while (i < 10) {
  /* 做点事情 */

  i++;
}
```

```c
int i = 0;

do {
  /* 做点事情 */

  i++;
} while (i < 10)
```

可以使用  break 跳出循环

```c
for (int i = 0; i <= 10; i++) {
  if (i == 4 && someVariable == 10) {
    break;
  }
}

int i = 0;
while (1) {
  /* 做点事情 */

  i++;
  if (i == 10) break;
}
```

# 数组

```c
int prices[5];

const int SIZE = 5;
int prices[SIZE];

int prices[5] = { 1, 2, 3, 4, 5 };

int prices[5];
prices[0] = 1;
prices[1] = 2;
prices[2] = 3;
prices[3] = 4;
prices[4] = 5;
```

**必须总是声明数组的大小。C 没有提供开箱即用的动态数组。**

C 数组中的所有元素都是顺序存放的，一个接一个。

**数组的变量名**，上述示例中的 prices，是**一个指向数组中首个元素的指针**。因此，可以像普通指针一样使用数组。

# 字符串

在 C 中，字符串是一种特殊的数组：**字符串是由 char 值组成的数组**：

```c
char name[7] = { "F", "l", "a", "v", "i", "o" };
char name[7] = "Flavio";
```

注意到“Flavio”是 6 个字符长，但是我定义了一个长度为 7 的数组吗？这是因为**字符串中的最后一个字符必须是 0（零）**，它是字符串的**终止符号**，我们必须给它留个位置。

```c
#include <string.h>
```

- strcpy()：将一个字符串复制到另一个字符串
- strcat()：将一个字符串追加到另一个字符串
- strcmp()：比较两个字符串是否相等
- strncmp()：比较两个字符串的前 n 个字符
- strlen()：计算字符串的长度

# 指针

**指针是某个内存块的地址，这个内存块包含一个变量。**

当你像这样声明一个整数时：

```c
int age = 37;
```

我们可以使用 & 运算符获取内存中该变量的地址值：

```c
printf("%p", &age); /* 0x7ffeef7dcb9c */
```

我们可以将该地址赋给一个变量：

```c
int address = &age;
```

当在声明中使用 int *address 时，我们并没有在声明一个整数值，而是在声明一个 **指向一个整数的指针**。

我们可以使用指针运算符获取该地址指向的变量的值：

```c
int age = 37;
int *address = &age;
printf("%u", *address); /* 37 */
```

我们又一次使用指针运算符，但是由于这次它不是一个声明，所以它表示“**该指针指向的变量的值**”。

# 函数

```c
void doSomething(int value) {
    printf("%u", value);
}
```

函数有 4 个重要的方面：

1. 它们有一个名字，所以我们可以在之后调用它们
2. 它们声明一个返回值
3. 它们可以有参数
4. 它们有一个函数体，用花括号包裹

如果函数没有返回值，你可以在函数名前面使用关键字 void。否则你就要声明该函数的返回值类型（整数为 int，浮点数为  float，字符串为 const char *，等等）。

函数返回值的数量不能超过一个。

函数可以有参数。它们是可选的。如果函数没有参数，我们就在括号内插入 void。

参数是通过 **拷贝** 传递的。这意味着如果你修改 value1，它的值是在局部作用域内修改的。函数外的那个值，即我们在调用时传入的值，并不会改变。

如果你传入的参数为一个 **指针**，你可以修改该变量的值，因为你现在可以使用它的内存地址直接访问它。

# 输入输出

在 C 中，我们有三种类型的 I/O 流：

- stdin（标准输入）
- stdout（标准输出）
- stderr（标准错误）

借助 I/O 函数，我们始终可以和流一起工作。流是一个高级接口，可以代表一个设备或文件。从 C 的角度来看，我们在从文件读取和命令行读取没有任何差异：不论如何，它都是一个 I/O 流。

某些函数是为与特定的流一起工作而设计的，就像 printf()一样，我们用它来将字符串打印到 stdout。使用它更加通用的版本 fprintf() 时，我们可以指定我们要写到的流。

还有其它像 %d 一样的格式指示符：

- %c 用于字符
- %s 用于字符串
- %f 用于浮点数
- %p 用于指针

```c
#include <stdio.h>

int main(void) {
  char name[20];
  printf("Enter your name: ");
  scanf("%s", name);
  printf("you entered %s", name);
}
```

# 变量作用域

当你在 C 程序中定义一个变量时，根据你声明它的位置，它会有一个不同的 **作用域（scope）**。

该位置决定了两种类型的变量：

- **全局变量（global variables）**
- **局部变量（local variables）**

在函数内部声明的变量就是局部变量。

局部变量只有在函数内才能访问，它们会在函数结束后不复存在。它们会被从内存中清除掉（有一些例外）。

原因是局部变量默认是在 **栈（stack）** 上声明的，除非你使用指针在堆中显式地分配它们。但是这样一来，你就不得不自己管理内存了。

定义在函数外部的变量就是全局变量。

全局变量可以从程序中的任何一个函数访问，它们在整个程序的执行过程中都是可用的，直到程序结束。

## 静态变量

在函数内部，你可以使用 static 关键字初始化一个 **静态变量（static variable）**。

我说了“在函数内部”，因为全局变量默认就是静态的，所以没有必要再添加这个关键字。

**静态变量在没有声明初始值的时候会被初始化为 0，并且它会在函数调用中保持该值**。

```c
int incrementAge() {
  int age = 0;
  age++;
  return age;
}
```

如果我们调用一次 incrementAge()，我们将会得到返回值 1。如果我们再调用一次，我们总是会得到 1，因为 age 是一个局部变量并且在每次调用该函数的时候都会被重新初始化为 0。

```c
int incrementAge() {
  static int age = 0;
  age++;
  return age;
}
```

现在我们每调用一次这个函数，我们就会得到一个增加了的值。

我们也可以有静态数组。这时，每一个数组元素都被初始化为 0：

## 全局变量

全局变量可以被程序内的任何函数访问。该访问并不只局限于读取全局变量的值：**任何函数都可以更新全局变量的值**。

因此，全局变量是一种在函数间共享相同数据的一种方式。

局部变量的主要不同在于，分配给局部变量的内存会在函数结束之后立即释放。

全局变量只在程序结束时才会释放。

# 类型定义

C 中的 **typedef** 关键字允许你定义新的类型。

按照惯例，我们创建的新类型通常是大写的。这样可以更加容易区分它，并且可以立即识别出它是一种类型。

```c
typedef existingtype NEWTYPE
```

# 枚举类型

使用 typedef 和 enum 关键字，我们可以定义具有指定值的类型。

```c
typedef enum {
  //值……
}
```

```c
#include <stdio.h>

typedef enum {
  monday,  
  tuesday,
  wednesday,
  thursday,
  friday,
  saturday,
  sunday
} WEEKDAY;

int main(void) {
  WEEKDAY day = monday;

  if (day == monday) {
    printf("It's monday!"); 
  } else {
    printf("It's not monday"); 
  }
}
```

枚举定义中的每个枚举项在内部都与一个整数配对。所以在这个示例中 monday 是 0，tuesday 是 1，以此类推。

这意味着对应的条件可以是 if (day == 0) 而不是 if (day == monday)，但是对于我们人类来说，使用名字比数字更合理，所以它是一个非常便利的语法。

# 结构体

利用 **struct** 关键字，我们可以使用基本的 C 类型创建复杂的数据结构。

结构体是一组由不同类型的值组成的集合。C 中的数组被限制为一种类型，所以结构体在很多用例中会显得非常有趣。

```c
struct <structname> {
  //变量……
};
```

```c
struct person {
  int age;
  char *name;
};
```

通过将变量添加到右花括号之后，分号之前，你可以声明类型为该结构体的变量，就像这样：

```c
struct person {
  int age;
  char *name;
} flavio;
struct person {
  int age;
  char *name;
} flavio, people[20];
```

声明一个名为 flavio 的 person 变量，以及一个具有 20 个 person 的名为 people 的数组。

```c
struct person {
  int age;
  char *name;
};

struct person flavio;
struct person flavio1 = { 37, "Flavio" };
```

一旦定义了结构体，我们就可以使用一个点（ . ）来访问它里面的值了：

```c
struct person {
  int age;
  char *name;
};

struct person flavio = { 37, "Flavio" };
printf("%s, age %u", flavio.name, flavio.age);
```

我们也可以使用点语法改变结构体中的值：

```c
struct person {
  int age;
  char *name;
};

struct person flavio = { 37, "Flavio" };

flavio.age = 38;
```

注意到结构体是 **复制传递** 的，这一点很重要，除非，当然你可以传递一个指向结构体的指针，这种情况下它就是引用传递。

使用 typedef，我们可以简化处理结构体时的代码。

```c
typedef struct {
  int age;
  char *name;
} PERSON;
PERSON flavio;
PERSON flavio = { 37, "Flavio" }; // 对比上面的，简化了 struct 关键字
```

# 命令行参数

在 C 程序中，你可能需要在命令启动时从命令行接收参数。

```c
int main(void)
int main (int argc, char *argv[])
```

argc 是一个整数，包含从命令行提供的参数的数量。

argv 是一个字符串数组。

当程序开始运行时，我们用这两个参数给主函数提供参数。

注意 argv 数组中总是至少有一个元素：程序的名字。

咱们写一个打印它收到的参数的程序吧：

```c
#include <stdio.h>

int main (int argc, char *argv[]) {
  for (int i = 0; i < argc; i++) {
    printf("%s\n", argv[i]);
  }
}
```

如果我们传递一些随机参数，就像这样：./hello a b c，我们竟会在终端中得到这个输出：

```
./hello
a
b
c
```

对于简单的需求而言，这个系统工作得很好。对于更加复杂的需求，有一些常用的包，比如 **getopt**。

# 头文件

简单的程序可以直接放在单个文件中。但是当你的程序变大，将它放在单个文件中就不可能了。

你可以将程序一些部分移动到一个单独的文件中，然后创建一个 **头文件**。

头文件看起来就像普通的 C 文件一样，但是它是以 .h 而不是 .c 结尾的。它里面的内容是 **声明**，而不是函数的实现和程序的其它部分。

```c
#include <stdio.h>
```

`#include` 是一个预处理器指令。

该预处理器会在标准库中寻找 stdio.h 文件，因为你使用了花括号包裹它。若要包含你自己的头文件，你需要使用引号（ " ），就像这样：

```c
#include "myfile.h"
```

上述代码会让预处理器在**当前文件夹内**寻找 myfile.h。

你也可以使用文件夹结构的库：

```c
#include "myfolder/myfile.h"
```

```c
#include <stdio.h>

int calculateAge(int year) {
  const int CURRENT_YEAR = 2020;
  return CURRENT_YEAR - year;
}

int main(void) {
  printf("%u", calculateAge(1983));
}
```

假设我想将 caculateAge 函数移到一个单独的文件中。

我创建一个名为 calculate_age.c 的文件：

```c
int calculateAge(int year) {
  const int CURRENT_YEAR = 2020;
  return CURRENT_YEAR - year;
}
```

我还创建了一个名为 calculate_age.h 的文件，我在其中放入了 **函数原型**，除了函数体，它与 .c 文件中的函数完全相同：

```c
int calculateAge(int year);
```

现在在主 .c 文件中，我们可以移除 calculateAge() 函数的定义，并且我们可以导入 calculate_age.h，它会让 calculateAge() 函数可用：

```c
#include <stdio.h>
#include "calculate_age.h"

int main(void) {
  printf("%u", calculateAge(1983));
}
```

别忘了编译多个文件组成的程序，你需要在命令行中列出它们，就像这样

```sh
gcc -o main main.c calculate_age.c
```

如果配置更加复杂，一个告诉编译器如何编译该程序的 Makefile 是必需的。

# 预处理器

预处理器是一个工具，当我们用 C 编程时，它对我们有很大的帮助。它是 C 标准的一部分，就像语言本身、编译器和标准库一样。

它解析我们的程序，确保编译器在处理之前获得所有需要的东西。

在实践中，它是做什么的呢？

例如，它查找你使用 #include 指令包含的所有头文件。

它还查看你使用 #define 定义的每个常量并将其替换为实际的值。

## 条件

我们能做的一件事情是使用条件让表达式决定程序的编译方式。

例如，我们可以检查 DEBUG 常量是否为 0：

```c
#include <stdio.h>

const int DEBUG = 0;

int main(void) {
#if DEBUG == 0
  printf("I am NOT debugging\n");
#else
  printf("I am debugging\n");
#endif
}
```

## 符号常量

```c
#define VALUE 1
#define PI 3.14
#define NAME "Flavio"
```

## 宏

宏与符号常量之间的差别在于：宏可以接受一个参数，并且通常包含代码，而符号常量只是一个值：

```c
#define POWER(x) ((x) * (x))
```

注意参数两侧的括号：当宏在预编译过程中被替换时，这是一个避免问题的好方法。

然后我们可以在代码中使用它，像这样：

```c
printf("%u\n", POWER(4)); //16
```

它与函数之间的一个大差别就是：**宏不会声明参数或返回值的类型**，这在一些场景中可能很方便。

然而，**宏的定义被限制成只有一行**。

## If defined

我们可以使用  #ifdef 来检查某个符号常量或宏是否被定义过：

```c
#include <stdio.h>
#define VALUE 1

int main(void) {
#ifdef VALUE
  printf("Value is defined\n");
#else
  printf("Value is not defined\n");
#endif
}
```

我们也可以使用 `#ifndev` 检查对立面（宏未定义）。

我们还可以使用 `#if defined` 和 `#if !defined` 来达到同样的目的。

像这样将一些代码块包裹到单个块中是很常见的：

```c
#if 0

#endif
```

这样可以临时防止程序运行，也可以使用一个 DEBUG 符号常量

```c
#define DEBUG 0

#if DEBUG
  // 当 DEBUG 不为 0 时，代码才会被发给编译器
#endif
```

## 你可以使用的预定义的符号常量

预处理器还定义了很多你可以直接使用的符号常量，它们的名字的前后有两个下划线作为标识，包括：

- **__LINE__** 代表源代码文件中的当前行
- **__FILE__** 代表文件的名字
- **__DATE__** 表示编译日期，格式为 Mmm gg aaaa
- **__TIME__** 表示编译实践，格式为 hh:mm:ss