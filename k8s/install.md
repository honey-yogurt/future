

ubuntu 安装 k8s 过程记录，详情参考[官方教程](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/)



## 集群配置

三台 ubuntu 机器，ip 分别为:

+ 192.168.1.115     root  qwer@123
+ 192.168.3.51       root  zxt123456
+ 192.168.3.155     root zkjg@123



## ubuntu 设置

### 配置ubuntu系统国内源

+ 备份默认源：

```bash
sudo cp /etc/apt/sources.list /etc/apt/sources.list.bak
sudo rm -rf /etc/apt/sources.list
```

+ 配置国内源（清华）：

```bash
sudo vim /etc/apt/sources.list
```

内容如下：

```text
# 默认注释了源码镜像以提高 apt update 速度，如有需要可自行取消注释
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-updates main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-backports main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-backports main restricted universe multiverse

# deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-security main restricted universe multiverse
# # deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-security main restricted universe multiverse

deb http://security.ubuntu.com/ubuntu/ jammy-security main restricted universe multiverse
# deb-src http://security.ubuntu.com/ubuntu/ jammy-security main restricted universe multiverse

# 预发布软件源，不建议启用
# deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-proposed main restricted universe multiverse
# # deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-proposed main restricted universe multiverse
```

+ 更新

修改完毕后，需要执行以下命令生效

```bash
sudo apt-get update
sudo apt-get upgrade
```

### 设置主机名



在master节点

```bash
sudo hostnamectl set-hostname master
```

在另外两个节点

```bash
sudo hostnamectl set-hostname work1
sudo hostnamectl set-hostname work2
```

在大型集群中，可以通过自动化工具来设置，如：Ansible、Puppet、Chef等

在Kubernetes中，主机名的设置对于集群的正常运行和通信非常重要，它影响到以下几个方面：

1. 节点识别和通信：Kubernetes集群由多个节点组成，每个节点都需要一个唯一的标识符来进行识别和通信。设置主机名可以确保每个节点具有唯一的标识符，这样Kubernetes控制平面和其他节点就能够正确地识别和与之通信。
2. DNS解析：Kubernetes使用DNS服务来解析集群内部的服务和Pod的域名。设置主机名可以确保节点的名称能够正确地解析为相应的IP地址，这样其他节点和服务就能够通过主机名进行通信。
3. SSL证书：Kubernetes使用SSL证书来加密和保护集群内部的通信。SSL证书通常使用主机名作为标识符，用于验证通信的双方身份。如果主机名没有正确设置，可能会导致证书验证失败，从而导致通信故障或安全问题

 ### 设置 hosts

方便节点之间相互访问，在每个节点的 /etc/hosts 上添加如下配置：

```tex
192.168.1.115 master
192.168.3.155 work1
192.168.3.51 work2
```



在Kubernetes集群中，设置hosts文件是为了实现节点之间的主机名解析。虽然Kubernetes集群中通常会使用DNS来进行服务发现和解析，但在某些情况下，直接使用hosts文件可以提供更快速和可靠的解析方式。

以下是设置hosts文件的几个原因：

1. 解析性能：使用hosts文件可以避免DNS查询的延迟，因为hosts文件是本地文件，解析速度更快。对于一些对解析速度要求较高的场景，使用hosts文件可以提供更低的延迟。
2. DNS故障：在某些情况下，DNS服务可能会发生故障或不可用。如果集群中的节点依赖于DNS服务进行主机名解析，那么当DNS不可用时，节点之间的通信可能会受到影响。通过设置hosts文件，可以在DNS故障时提供备用的解析方式，确保节点之间的通信正常进行。
3. 高可用性：在某些情况下，当节点无法通过DNS解析来识别其他节点时，使用hosts文件可以提供一种备用的识别方式。这对于某些关键组件（如控制平面节点）的高可用性非常重要，因为它们需要快速准确地识别其他节点。

### **禁用** selinux

默认ubuntu下没有这个模块，centos下需要禁用selinux；

SELinux 是一个强制访问控制（MAC）机制，用于加强 Linux 系统的安全性。它通过对进程和文件系统的访问进行细粒度的控制，可以限制恶意软件或攻击者对系统的访问和操作。然而，由于 Kubernetes 需要在节点上运行多个容器和进程，并且需要与其他节点和控制平面进行通信，这可能会导致一些 SELinux 相关的问题：

1. 容器访问权限：SELinux 可能会限制容器对主机文件系统的访问权限。由于容器通常需要与主机共享一些资源，如日志文件、配置文件等，禁用 SELinux 可以避免容器无法访问所需资源的问题。
2. 容器网络通信：SELinux 可能会限制容器之间的网络通信。Kubernetes 中的容器通常需要通过网络进行通信，例如，Pod 中的不同容器之间需要相互通信，禁用 SELinux 可以避免网络通信受到限制。
3. 安装和配置问题：在某些情况下，SELinux 可能会干扰 Kubernetes 的安装和配置过程。例如，它可能会阻止某些必需的文件或进程的访问，导致安装或配置失败。

随着 Kubernetes 的不断发展和更新，一些与 SELinux 相关的问题可能已经得到解决，因此在特定情况下，可能不需要禁用 SELinux。

**避免在学习的第一步“安装”就出现未知错误，建议禁用，降低学习门槛。**

### 禁用 swap

 Kubernetes 对于节点上的内存管理和容器调度假设系统没有启用 swap。

Swap 是一种用于将内存中不常用的数据暂时存储到磁盘上的机制。当系统的物理内存不足时，操作系统将部分内存中的数据移动到 swap 空间，以释放内存供其他进程使用。然而，对于 Kubernetes 集群而言，禁用 swap 有以下几个原因：

1. 容器资源限制：Kubernetes 使用 cgroups（控制组）来对容器进行资源限制和管理。cgroups 可以限制容器使用的 CPU、内存等资源。然而，当系统启用 swap 时，cgroups 无法有效地限制容器使用的内存，因为容器可以将不常用的内存数据交换到 swap 空间，绕过了 cgroups 的限制。
2. 容器调度和性能：Kubernetes 调度器负责将容器调度到节点上，并根据节点的资源情况进行分配。当系统启用 swap 时，节点上的内存资源包括物理内存和 swap 空间，这可能会导致 Kubernetes 错误地将容器调度到资源不足的节点上，从而影响容器的性能和稳定性。
3. OOM（Out of Memory）事件处理：当系统的内存资源耗尽时，操作系统会触发 OOM 事件，选择终止某个进程以释放内存。当系统启用 swap 时，OOM 事件可能会导致 Kubernetes 节点上的容器被终止，从而影响应用程序的可用性和稳定性。

+ 临时禁用

```bash
sudo swapoff -a
```

+ 永久禁用

```bash
sudo vim /etc/fstab
```

将最后一行注释后重启系统即生效：

```bash
 #/swapfile                                 none            swap    sw              0      0
```

```bash
sudo reboot
swapon --show # 什么都不输出，即禁用成功
```



### 修改内核参数，转发 IPv4 并让 iptables 看到桥接流量

```bash
 sudo tee /etc/modules-load.d/containerd.conf <<EOF
 overlay
 br_netfilter
 EOF
 
 sudo modprobe overlay
 sudo modprobe br_netfilter

```

```bash
 sudo tee /etc/sysctl.d/kubernetes.conf <<EOF
 net.bridge.bridge-nf-call-ip6tables = 1
 net.bridge.bridge-nf-call-iptables = 1
 net.ipv4.ip_forward = 1
 EOF

```

```bash
sudo sysctl --system
```

加载模块：**overlay** 和 **br_netfilter** 是两个内核模块，它们在使用容器运行时引擎（例如 containerd）时需要加载。这些模块提供了一些必要的功能和驱动程序，以便容器可以在 Linux 上运行。

**net.bridge.bridge-nf-call-ip6tables** 和 **net.bridge.bridge-nf-call-iptables**：这两个参数控制桥接网络的网络地址转换（NAT）功能是否对 IPv6 和 IPv4 包进行调用。Kubernetes 使用网络地址转换来实现 Pod 和 Service 的网络通信。将这些参数设置为 1，允许桥接网络对通过 iptables 和 ip6tables 进行的网络流量进行 NAT 处理。

**net.ipv4.ip_forward** ：该参数启用 IP 转发功能，允许 Linux 内核将网络流量从一个网络接口转发到另一个网络接口。在 Kubernetes 集群中，它用于将流量从一个 Pod 转发到另一个 Pod 或外部网络。

通过运行以下指令确认 br_netfilter 和 overlay 模块被加载：

```bash
lsmod | grep br_netfilter
lsmod | grep overlay
```

通过运行以下指令确认 net.bridge.bridge-nf-call-iptables、net.bridge.bridge-nf-call-ip6tables 和 net.ipv4.ip_forward 系统变量在你的 sysctl 配置中被设置为 1：

```bash
sysctl net.bridge.bridge-nf-call-iptables net.bridge.bridge-nf-call-ip6tables net.ipv4.ip_forward
```



## 容器运行时

### 安装 containerd

安装dependencies，必需的依赖项

```bash
sudo apt install -y curl gnupg2 software-properties-common apt-transport-https ca-certificates

```





## kubeadm

### 简介

部署和管理 Kubernetes 集群：

+ 初始化集群：使用 kubeadm 可以方便地初始化一个全新的 Kubernetes 集群。它会自动配置并启动必要的组件，如 etcd、kube-apiserver、kube-controller-manager 和 kube-scheduler。
+ 加入节点：kubeadm 提供了加入节点到集群的功能。通过在新节点上运行 kubeadm join 命令，并提供集群的连接信息，新节点可以加入到现有的 Kubernetes 集群中。
+ 组件管理：kubeadm 可以帮助管理 Kubernetes 集群的组件。它可以升级和更新组件的版本，检查组件的健康状态，并提供一些诊断和故障排除的工具。
+ 配置管理：kubeadm 允许用户通过配置文件来定义和自定义集群的配置。可以使用 kubeadm init 阶段生成的配置文件，或者使用 kubeadm config 命令来生成和修改配置文件。
+ 安全性设置：kubeadm 提供了一些安全性设置的选项，以确保集群的安全性。例如，可以使用 kubeadm 进行证书管理，生成和轮换证书，以及设置 RBAC（Role-Based Access Control）规则。

### 安装



## kubelet

### 简介

kubelet 是 Kubernetes 集群中的一个核心组件，它在每个节点上运行，并负责管理节点上的容器和与控制平面的通信：

+ 容器生命周期管理：kubelet 负责监视节点上的容器，并确保它们按照所定义的期望状态运行。它会根据控制平面发送的 Pod 配置信息，创建、启动、停止和删除容器。如果容器失败或被删除，kubelet 会尝试重新启动或清理它们。
+ 资源管理：kubelet 监视节点上的资源使用情况，并根据集群的资源分配策略来管理资源。它会根据容器的资源需求和节点的资源容量来调度和分配资源，以确保各个容器能够正常运行且不超出节点的资源限制。
+ 网络管理：kubelet 负责为容器配置网络。它会与容器运行时（如 Docker 或 containerd）交互，为容器分配 IP 地址，并设置容器网络的路由和防火墙规则。此外，kubelet 还与网络插件协作，确保容器可以与其他容器和服务进行通信。
+ 健康检查和自愈：kubelet 定期检查运行在节点上的容器的健康状态。如果容器出现故障或不响应，kubelet 会尝试重新启动容器，以使其恢复正常。如果多次重启失败，kubelet 会将容器标记为失败，并通知控制平面。
+ 日志和监控：kubelet 收集容器的日志和监控信息，并将其发送到集群的日志和监控系统中。这样，用户可以方便地查看和分析容器的运行情况，以及进行故障排除和性能优化。

### 安装

## kubectl

### 简介

kubectl 是 Kubernetes 的命令行工具，用于与 Kubernetes 集群进行交互和管理：

+ 配置和管理资源：kubectl 可以使用 YAML 或 JSON 格式的配置文件来创建、更新和删除 Kubernetes 集群中的资源对象，如 Pod、Deployment、Service、ConfigMap 等。它可以通过执行命令来创建或修改这些资源，并将其发送到集群的 API 服务器进行处理。
+ 查看集群状态：kubectl 提供了一系列命令，用于查看集群中各种资源的状态和信息。例如，可以使用 kubectl get 命令来列出集群中的资源对象，使用 kubectl describe 命令来获取资源的详细信息，使用 kubectl logs 命令来查看容器的日志等。
+ 调试和故障排除：kubectl 提供了一些调试和故障排除的功能。例如，可以使用 kubectl exec 命令在容器内部执行命令，以检查容器的运行环境和状态。还可以使用 kubectl logs 命令查看容器的日志，以帮助分析和解决问题。
+ 执行操作和管理任务：kubectl 允许执行一些操作和管理任务，如扩容和缩容应用程序、滚动更新应用程序、暂停和恢复应用程序的部署、管理存储卷等。可以使用 kubectl scale、kubectl rollout、kubectl pause 和 kubectl resume 等命令来执行这些任务。
+ 访问集群和控制平面：kubectl 可以配置和管理与 Kubernetes 集群的连接。它可以使用配置文件或命令行参数指定集群的连接信息，包括 API 服务器的地址和凭据。通过 kubectl，可以与集群进行交互，并执行各种管理操作。

### 安装

1. Download the latest release with the command:

```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```

2. Install kubectl

```bash
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

3. Test to ensure the version you installed is up-to-date:

```bash
kubectl version --client --output=yaml
```

