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
chgrp: invalid group: `testing' &lt;== 发生错误，找不到这个群组名
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



## Linux 目录配置

| 一级目录 | 功能（作用）                                                 |
| :------- | :----------------------------------------------------------- |
| /bin/    | 存放系统命令，普通用户和 root 都可以执行。放在 /bin 下的命令在单用户模式下也可以执行 |
| /boot/   | 系统启动目录，保存与系统启动相关的文件，如内核文件和启动引导程序（grub）文件等 |
| /dev/    | 设备文件保存位置                                             |
| /etc/    | **配置文件保存位置**。系统内所有采用默认安装方式（rpm 安装）的服务配置文件全部保存在此目录中，如用户信息、服务的启动脚本、常用服务的配置文件等 |
| /home/   | 普通用户的主目录（也称为家目录）。在创建用户时，每个用户要有一个默认登录和保存自己数据的位置，就是用户的主目录，所有普通用户的主目录是在 /home/ 下建立一个和用户名相同的目录。如用户 liming 的主目录就是 /home/liming |
| /lib/    | 系统调用的函数库保存位置                                     |
| /media/  | 挂载目录。系统建议用来挂载媒体设备，如软盘和光盘             |
| /mnt/    | 挂载目录。早期 Linux 中只有这一个挂载目录，并没有细分。系统建议这个目录用来挂载额外的设备，如 U 盘、移动硬盘和其他操作系统的分区 |
| /misc/   | 挂载目录。系统建议用来挂载 NFS 服务的共享目录。虽然系统准备了三个默认挂载目录 /media/、/mnt/、/misc/，但是到底在哪个目录中挂载什么设备可以由管理员自己决定。 |
| /opt/    | 第三方安装的软件保存位置。这个目录是放置和安装其他软件的位置，手工安装的源码包软件都可以安装到这个目录中。 |
| /root/   | root 的主目录。普通用户主目录在 /home/ 下，root 主目录直接在“/”下 |
| /sbin/   | 保存与系统环境设置相关的命令，只有 root 可以使用这些命令进行系统环境设置，但也有些命令可以允许普通用户查看 |
| /srv/    | 服务数据目录。一些系统服务启动之后，可以在这个目录中保存所需要的数据 |
| /tmp/    | **临时目录**。系统存放临时文件的目录，在该目录下，所有用户都可以访问和写入。建议此目录中不能保存重要数据，最好每次开机都把该目录清空 |

FHS 针对根目录中包含的子目录仅限于上表，但除此之外，Linux 系统根目录下通常还包含下表中的几个一级目录。

| 一级目录     |                         功能（作用）                         |
| :----------- | :----------------------------------------------------------: |
| /lost+found/ | 当系统意外崩溃或意外关机时，产生的一些文件碎片会存放在这里。在系统启动的过程中，fsck 工具会检查这里，并修复已经损坏的文件系统。这个目录只在每个分区中出现，例如，/lost+found 就是根分区的备份恢复目录，/boot/lost+found 就是 /boot 分区的备份恢复目录 |
| /proc/       | 虚拟文件系统。**该目录中的数据并不保存在硬盘上，而是保存到内存中**。**主要保存系统的内核、进程、外部设备状态和网络状态等**。如 /proc/cpuinfo 是保存 CPU 信息的，/proc/devices 是保存设备驱动的列表的，/proc/filesystems 是保存文件系统列表的，/proc/net 是保存网络协议信息的...... |
| /sys/        | 虚拟文件系统。和 /proc/ 目录相似，该目录中的数据都保存在内存中，主要保存与内核相关的信息 |



usr（注意不是 user），全称为 Unix Software Resource，此目录用于存储**系统软件资源**。FHS 建议所有开发者，应把软件产品的数据合理的放置在 /usr 目录下的各子目录中，而不是为他们的产品创建单独的目录。

Linux 系统中，所有系统默认的软件都存储在 /usr 目录下

FHS 建议，/usr 目录应具备下表所示的子目录。

| 子目录       | 功能（作用）                                                 |
| :----------- | :----------------------------------------------------------- |
| /usr/bin/    | 存放系统命令，普通用户和超级用户都可以执行。这些命令和系统启动无关，在单用户模式下不能执行 |
| /usr/sbin/   | 存放根文件系统不必要的系统管理命令，如多数服务程序，只有 root 可以使用。 |
| /usr/lib/    | 应用程序调用的函数库保存位置                                 |
| /usr/XllR6/  | 图形界面系统保存位置                                         |
| /usr/local/  | 手工安装的软件保存位置。我们一般建议源码包软件安装在这个位置 |
| /usr/share/  | 应用程序的资源文件保存位置，如帮助文档、说明文档和字体目录   |
| /usr/src/    | 源码包保存位置。我们手工下载的源码包和内核源码包都可以保存到这里。 |
| /usr/include | C/C++ 等编程语言头文件的放置目录                             |

> 在单用户模式下，系统启动时只加载最基本的组件和服务，不加载用户级别的环境和配置。单用户模式是用于系统故障修复和维护的模式，只有超级用户（root）可以登录并执行必要的操作。
>
> 在单用户模式下，系统处于最小化状态，仅提供最基本的功能和服务。这是为了确保系统的稳定性和安全性，避免执行不必要的命令和操作可能引发的问题。
>
> 普通用户的环境和配置文件通常包含了许多自定义设置和路径，其中可能包含与系统命令相关的信息。在单用户模式下，这些环境和配置文件不会被加载，因此普通用户无法执行系统命令。





/var 目录用于**存储动态数据**，例如缓存、日志文件、软件运行过程中产生的文件等。通常，此目录下建议包含如下表所示的这些子目录。

| /var子目录        | 功能（作用）                                                 |
| :---------------- | :----------------------------------------------------------- |
| /var/lib/         | 程序运行中需要调用或改变的数据保存位置。如 MySQL 的数据库保存在 /var/lib/mysql/ 目录中 |
| /var/log/         | 登陆文件放置的目录，其中所包含比较重要的文件如 /var/log/messages, /var/log/wtmp 等。 |
| /var/run/         | 一些服务和程序运行后，它们的 PID（进程 ID）保存位置          |
| /var/spool/       | 里面主要都是一些临时存放，随时会被用户所调用的数据，例如 /var/spool/mail/ 存放新收到的邮件，/var/spool/cron/ 存放系统定时任务。 |
| /var/www/         | RPM 包安装的 Apache 的网页主目录                             |
| /var/nis和/var/yp | NIS 服务机制所使用的目录，nis 主要记录所有网络中每一个 client 的连接信息；yp 是 linux 的 nis 服务的日志文件存放的目录 |
| /var/tmp          | 一些应用程序在安装或执行时，需要在重启后使用的某些文件，此目录能将该类文件暂时存放起来，完成后再行删除 |



## 目录的相关操作

```bash
. 代表此层目录
.. 代表上一层目录
- 代表前一个工作目录
~ 代表“目前使用者身份”所在的主文件夹
~account 代表 account 这个使用者的主文件夹（account是个帐号名称）

```

根目录下确实存在 . 与 .. 两个目录， 可发现这两个目录的属性与权限完全一致，这代表根目录的上一层（..）与根目录自己（.）是同一个目录。



## 关于可执行文件路径的变量： $PATH

当我们在执行一个指令的时候，举例来说“ls”好了，系统会依照PATH的设置去每个PATH定义 的目录下搜寻文件名为ls的可执行文件， 如果在PATH定义的目录中含有多个文件名为ls的可执行文件，那么先搜寻到的同名指令先被执行！