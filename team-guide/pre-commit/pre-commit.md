# Git hook

Git 能在特定的重要动作发生时触发自定义脚本钩子。钩子分为两组：

- 客户端钩子：pre-commit, prepare-commit-msg, commit-msg, post-commit 等，主要在服务端接收提交对象时、推送到服务器之前调用。
- 服务器钩子：pre-receive, post-receive, update 等，主要在服务端接收提交对象时、推送到服务器之前调用。

# Introduction

Git hook scripts 在提交代码审查之前识别简单问题非常有用。

随着我们创建了更多的库和项目，我们意识到跨项目共享预提交钩子是很麻烦的。我们从一个项目复制并粘贴笨重的bash脚本到另一个项目，并且需要手动修改这些钩子以适应不同的项目结构。

我们认为您应该始终使用最佳行业标准规范检查工具。一些最好的规范检查工具是用您所使用项目中未安装或未使用语言编写的。例如scss-lint是一个用Ruby编写的SCSS规范检查工具。如果您正在node中编写一个项目，则应该能够将scss-lint作为预提交钩子使用，而无需向您的项目添加Gemfile文件或理解如何安装scss-lint。

pre-commit 是一个多语言包管理器，用于管理预提交钩子。您可以指定要使用哪些钩子，pre-commit会在每次提交之前安装和执行任何语言编写的钩子。pre-commit专门设计为不需要root访问权限。

# Installation

使用 pip

```
pip install pre-commit
```

