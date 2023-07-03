# Go Env

## 下载包

https://go.dev/doc/install

## 配置环境变量

### Remove any previous Go installation

```bash
 rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
```

### Add /usr/local/go/bin to the `PATH` environment variable

**vim /etc/profile**  or  **vim $HOME/.profile**

```
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

**source /etc/profile** or **source $HOME/.profile**

## go verson

```
go version
```



## 配置代理

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

