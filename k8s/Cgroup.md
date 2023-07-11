# 什么是 Cgroup

Cgroups （Control Groups）是 Linux 下用于对一个或一组进程进行资源控制和监控的机制；

+ 可以对诸如 CPU 使用时间、内存、磁盘 I/O 等进程所需的资源进行限制；
+ 不同资源的具体管理工作由相应的 Cgroup 子系统（Subsystem）来实现 ；
+ 针对不同类型的资源限制，只要将限制策略在不同的的子系统上进行关联即可 ；
+ Cgroups 在不同的系统资源管理子系统中以层级树（Hierarchy）的方式来组织管理：每个 Cgroup 都可以 包含其他的子 Cgroup，因此子 Cgroup 能使用的资源除了受本 Cgroup 配置的资源参数限制，还受到父 Cgroup 设置的资源限制 。



# Linux 内核代码中 Cgroups 的实现

进程数据结构

```c 
struct task_struct
{
	#ifdef CONFIG_CGROUPS
	struct css_set __rcu *cgroups;
	struct list_head cg_list;
	#endif
}
```

css_set 是 cgroup_subsys_state 对象的集合数据结构

```c
struct css_set {
/*
* Set of subsystem states, one for each subsystem. This array is
* immutable after creation apart from the init_css_set during
* subsystem registration (at boot time).
*/
	struct cgroup_subsys_state *subsys[CGROUP_SUBSYS_COUNT];
};
```

