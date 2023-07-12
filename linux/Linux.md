# 文件与目录管理

## 使用者与群组

Linux 一般将文件可存取的身份分为三个类别，分别是 owner/group/others，且三种身份各有 read/write/execute 等权限。

Linux 系统当中，默认的情况下，所有的系统上的帐号与一般身份使用者，还有 root 的相关信息， 都是记在 /etc/passwd这个文件内的。至于个人的密码则是记录在 /etc/shadow 这个文件下。 此外，Linux所有的群组名称都纪录在/etc/group内！

## 文件权限

```bash
ls -al
-rw-r--r--  1 root root     3106 10月 15  2021 .bashrc
drwxr-xr-x  3 root root     4096  6月 12 15:32 go
```

+ 第一栏代表这个文件的类型与权限（permission）

第一栏共有 10 个字符：

1. 第一个字符代表这个文件是“目录、文件或链接文件等等”：
   1. 当为[ d ]则是目录
   2. 当为[ - ]则是文件
   3. 若是[ l ]则表示为链接文件
   4. 若是[ b ]则表示为设备文件里面的可供储存的周边设备（可随机存取设备）
   5. 若是[ c ]则表示为设备文件里面的串行端口设备，例如键盘、鼠标（一次性读取设备）

2. 接下来的字符中，以三个为一组，且均为“rwx” 的三个参数的组合。其中，[ r ]代表可读 （read）、[ w ]代表可写（write）、[ x ]代表可执行（execute），如果是目录的话，没有[ x ]不能进入该目录。 要注意的是，这三个权限的位置不会改变，如果没有权限，就会出现减号[ - ]而已。
   1. 第一组为“文件拥有者可具备的权限”
   2. 第二组为“加入此群组之帐号的权限”
   3. 第三组为“非本人且没有加入本群组之其他帐号的权限”

+ 第二栏表示有多少文件名链接到此节点（i-node）：

每个文件都会将他的权限与属性记录到文件系统的i-node中，不过，我们使用的目录树却是使用文件名来记录， 因此每个文件名就会链接到一个i-node。这个属性记录的，就是有多少不同的文件名链接到相同的一个i-node 。

+ 第三栏表示这个文件（或目录）的“拥有者帐号”
+ 第四栏表示这个文件的所属群组
+ 第五栏为这个文件的容量大小，默认单位为Bytes；
+ 第六栏为这个文件的创建日期或者是最近的修改日期
+ 第七栏为这个文件的文件名，如果文件名之前多一个“ . ”，则代表这个文件为“隐藏 文件”

## 改变文件权限

+ chgrp ：改变文件所属群组 
+ chown ：改变文件拥有者 
+ chmod ：改变文件的权限, SUID, SGID, SBIT等等的特性

chgrp : 要被改变的群组名称必须要 在 /etc/group 文件内存在才行，否则就会显示错误！

```bash
[root@study ~]# chgrp [-R] 群组 dirname/filename ...
选项与参数：
-R : 进行递回（recursive）的持续变更，亦即连同次目录下的所有文件、目录
都更新成为这个群组之意。常常用在变更某一目录内所有的文件之情况。
范例：
[root@study ~]# chgrp users initial-setup-ks.cfg
[root@study ~]# ls -l
-rw-r--r--. 1 root users 1864 May 4 18:01 initial-setup-ks.cfg
[root@study ~]# chgrp testing initial-setup-ks.cfg
chgrp: invalid group: `testing' &lt;== 发生错误 ～找不到这个群组名～
```

chown : 使用者必须是已经存在系统中的帐号，也就是在 /etc/passwd 这个文件中有纪录的使用者名称才能改变

```bash
[root@study ~]# chown [-R] 帐号名称 文件或目录
[root@study ~]# chown [-R] 帐号名称:群组名称 文件或目录
选项与参数：
-R : 进行递回（recursive）的持续变更，亦即连同次目录下的所有文件都变更
范例：将 initial-setup-ks.cfg 的拥有者改为bin这个帐号：
[root@study ~]# chown bin initial-setup-ks.cfg
[root@study ~]# ls -l
-rw-r--r--. 1 bin users 1864 May 4 18:01 initial-setup-ks.cfg
范例：将 initial-setup-ks.cfg 的拥有者与群组改回为root：
[root@study ~]# chown root:root initial-setup-ks.cfg
[root@study ~]# ls -l
-rw-r--r--. 1 root root 1864 May 4 18:01 initial-setup-ks.cfg
```

chmod : 文件权限的改变使用的是 chmod 这个指令，权限的设置方法有两种， 分别可以使用数 字或者是符号来进行权限的变更。  

+ 数字类型改变文件权限

```bash
r:4 > w:2 > x:1
```

每种身份（owner/group/others）各自的三个权限（r/w/x）分数是需要累加的，例如当权 限为： [-rwxrwx---] 分数则是：

```bash
owner = rwx = 4+2+1 = 7 > group = rwx = 4+2+1 = 7 > others= --- = 0+0+0 = 0
```

```bash
[root@study ~]# chmod [-R] xyz 文件或目录
选项与参数：
xyz : 就是刚刚提到的数字类型的权限属性，为 rwx 属性数值的相加。
-R : 进行递回（recursive）的持续变更，亦即连同次目录下的所有文件都会变更
[root@study ~]# ls -al .bashrc
-rw-r--r--. 1 root root 176 Dec 29 2013 .bashrc
[root@study ~]# chmod 777 .bashrc
[root@study ~]# ls -al .bashrc
-rwxrwxrwx. 1 root root 176 Dec 29 2013 .bashrc
```

+ 符号类型改变文件权限

```bash
| chmod | u g o a | +（加入） -（除去） =（设置） | r w x | 文件或目录 |
```

假如我们要“设置”一个文件的权限成为“-rwxr-xr-x”时，基本上就是：

user （u）：具有可读、可写、可执行的权限；

group 与 others （g/o）：具有可读与执行的权限。 所以就是：

```bash
[root@study ~]# chmod u=rwx,go=rx .bashrc
# 注意喔！那个 u=rwx,go=rx 是连在一起的，中间并没有任何空白字符！
[root@study ~]# ls -al .bashrc
-rwxr-xr-x. 1 root root 176 Dec 29 2013 .bashrc

[root@study ~]# ls -al .bashrc
-rwxr-xr-x. 1 root root 176 Dec 29 2013 .bashrc
[root@study ~]# chmod a+w .bashrc
[root@study ~]# ls -al .bashrc
-rwxrwxrwx. 1 root root 176 Dec 29 2013 .bashrc

[root@study ~]# chmod a-x .bashrc
[root@study ~]# ls -al .bashrc
-rw-rw-rw-. 1 root root 176 Dec 29 2013 .bashrc
[root@study ~]# chmod 644 .bashrc 

```

## 目录和文件之权限意义

+ 权限对文件的重要性
  + r （read）：可读取此一文件的实际内容，如读取文本文件的文字内容等；
  + w （write）：可以编辑、新增或者是修改该文件的内容（但不含删除该文件）；
  + x （eXecute）：该文件具有可以被系统执行的权限。

+ 权限对目录的重要性
  + r （read contents in directory）：表示具有读取目录结构清单的权限，所以当你具有读取（r）一个目录的权限时，表示你可以**查询**该目录下的文件名数据。
  + w （modify contents of directory）：
    + 创建新的文件与目录；
    + 删除已经存在的文件与目录（不论该文件的权限为何！）
    + 将已存在的文件或目录进行更名；
    + 搬移该目录内的文件、目录位置。
  + x （access directory）：目录的x代表的是使用者**能否进入该目录**成为工作目录的用途！



