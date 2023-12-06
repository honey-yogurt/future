# Git hook

Git 能在特定的重要动作发生时触发自定义脚本钩子。钩子分为两组：

- 客户端钩子：pre-commit, prepare-commit-msg, commit-msg, post-commit 等，主要在服务端接收提交对象时、推送到服务器之前调用。
- 服务器钩子：pre-receive, post-receive, update 等，主要在服务端接收提交对象时、推送到服务器之前调用。

git hooks 位置位于每个 git 项目下的 `.git/hooks` 目录里，进去后会看到这些钩子的官方示例，都是以 `.sample` 结尾的文件，这些示例脚本是不会执行的，去掉 `.sample` 后缀可激活该钩子脚本。

> GIt hooks 的每个钩子的作用和说明，详细的以官方文档为准： https://git-scm.com/docs/githooks



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

查看版本

```
pre-commit --version
pre-commit 3.5.0
```



# Configuration

创建配置文件 `.pre-commit-config.yaml`

pre-commit 以 `.pre-commit-config.yaml` 文件作为默认的配置文件，在项目根目录执行如下命令生成简单的配置内容：

```
pre-commit sample-config > .pre-commit-config.yaml
```

```yaml
# See https://pre-commit.com for more information
# See https://pre-commit.com/hooks.html for more hooks
repos:
-   repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v3.2.0
    hooks:
    -   id: trailing-whitespace
    -   id: end-of-file-fixer
    -   id: check-yaml
    -   id: check-added-large-files

```

- repos：表示一系列仓库的映射。
  - repo：表示接下来使用的 hooks 脚本从哪个仓库进行拉取。
  - rev：指定将要拉取的 tag 。
  - hooks：钩子脚本列表，这些脚本来自于 repo 定义的仓库中。
    - id：指定将要应用的钩子的名称，就是对应的文件名。

简单解释下如上配置文件的意思：在代码提交之前，会运行 hooks 列表中的这些检查，这些脚本来自于 `https://github.com/pre-commit/pre-commit-hooks` 这个仓库的 `v4.4.0`。四个检查脚本的含义如下：

- `trailing-whitespace` ：检查修建行尾的空格
- `end-of-file-fixer` ：确保文件以换行符结尾且仅以换行符结尾。
- `check-yaml` ：检查 yaml 文件的语法。
- `check-added-large-files` ：防止提交大文件。(默认检测阈值为 500KB)

the full set of options for the configuration are listed [here](https://pre-commit.com/#plugins)



# Install the git hook scripts

```sh
pre-commit install
```

现在pre-commit将自动运行于git commit！



# (optional) Run against all the files

通常在添加新的 hook 时，对所有文件运行钩子是一个好主意（通常 pre-commit 只会在 git 钩子期间运行更改的文件）。

```sh
re-commit run --all-files
```

```
[INFO] Initializing environment for https://github.com/pre-commit/pre-commit-hooks.
[INFO] Installing environment for https://github.com/pre-commit/pre-commit-hooks.
[INFO] Once installed this environment will be reused.
[INFO] This may take a few minutes...
Trim Trailing Whitespace.................................................Failed
- hook id: trailing-whitespace
- exit code: 1
- files were modified by this hook


```

可以看到会对我们所有的代码进行检查，并对不符合 check 的文件进行修改。



# Adding pre-commit plugins to your project

## Top Level

|                                                              |                                                              |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [`repos`](https://pre-commit.com/#top_level-repos)           | A list of [repository mappings](https://pre-commit.com/#pre-commit-configyaml---repos). |
| [`default_install_hook_types`](https://pre-commit.com/#top_level-default_install_hook_types) | （可选：默认[pre-commit]）一个--hook-types的列表，当运行pre-commit install时，默认使用它们。2.18.0中新增。 |
| [`default_language_version`](https://pre-commit.com/#top_level-default_language_version) | （可选：默认{}）从语言到应该用于该语言的默认language_version的映射。这只会覆盖没有设置language_version的个别钩子。 |
| [`default_stages`](https://pre-commit.com/#top_level-default_stages) | 可选的：默认（所有阶段））是钩子的 stages 属性的配置级别默认值。这只会覆盖那些没有设置 stages 的个别钩子。 |
| [`files`](https://pre-commit.com/#top_level-files)           | （可选：默认''）全局文件包含模式                             |
| [`exclude`](https://pre-commit.com/#top_level-exclude)       | （可选：默认^$）全局文件排除模式。                           |
| [`fail_fast`](https://pre-commit.com/#top_level-fail_fast)   | （可选：默认为false）设置为true以在第一个失败后停止运行pre-commit钩子。 |
| [`minimum_pre_commit_version`](https://pre-commit.com/#top_level-minimum_pre_commit_version) | （可选：默认为'0'）需要一个最低版本的pre-commit。            |



```
exclude: '^$'
fail_fast: false
repos:
-   ...
```

## repos 

|                                                |                                                              |
| ---------------------------------------------- | ------------------------------------------------------------ |
| [`repo`](https://pre-commit.com/#repos-repo)   | the repository url to `git clone` from                       |
| [`rev`](https://pre-commit.com/#repos-rev)     | the revision or tag to clone at.                             |
| [`hooks`](https://pre-commit.com/#repos-hooks) | A list of [hook mappings](https://pre-commit.com/#pre-commit-configyaml---hooks). |



## hooks

hook mappings 配置了使用仓库中的哪个钩子，并允许进行自定义。所有可选键将从仓库的配置中获取其默认值。

|                                                              |                                                              |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [`id`](https://pre-commit.com/#config-id)                    | which hook from the repository to use.<br />使用仓库中的哪个钩子。 |
| [`alias`](https://pre-commit.com/#config-alias)              | (optional) allows the hook to be referenced using an additional id when using `pre-commit run <hookid>`.<br />（可选）允许在使用 pre-commit run <hookid> 时，通过附加的 id 引用钩子。 |
| [`name`](https://pre-commit.com/#config-name)                | (optional) override the name of the hook - shown during hook execution.<br />（可选）覆盖挂钩的名称 - 在执行挂钩期间显示。 |
| [`language_version`](https://pre-commit.com/#config-language_version) | (optional) override the language version for the hook. See [Overriding Language Version](https://pre-commit.com/#overriding-language-version).<br />（可选）覆盖挂钩的语言版本。 |
| [`files`](https://pre-commit.com/#config-files)              | (optional) override the default pattern for files to run on.<br />（可选）覆盖默认的文件运行模式。 |
| [`exclude`](https://pre-commit.com/#config-exclude)          | (optional) file exclude pattern.<br />（可选）文件排除模式。 |
| [`types`](https://pre-commit.com/#config-types)              | (optional) override the default file types to run on (AND). See [Filtering files with types](https://pre-commit.com/#filtering-files-with-types).<br />（可选）覆盖默认的文件类型以运行（AND）。 |
| [`types_or`](https://pre-commit.com/#config-types_or)        | (optional) override the default file types to run on (OR). See [Filtering files with types](https://pre-commit.com/#filtering-files-with-types). *new in 2.9.0*.<br />（可选）覆盖默认文件类型以在（或者）上运行。 |
| [`exclude_types`](https://pre-commit.com/#config-exclude_types) | (optional) file types to exclude.<br />(可选) 要排除的文件类型。 |
| [`args`](https://pre-commit.com/#config-args)                | (optional) list of additional parameters to pass to the hook.<br />（可选）传递给钩子的附加参数列表。 |
| [`stages`](https://pre-commit.com/#config-stages)            | (optional) selects which git hook(s) to run for. See [Confining hooks to run at certain stages](https://pre-commit.com/#confining-hooks-to-run-at-certain-stages).<br />（可选）选择要运行的git钩子。 |
| [`additional_dependencies`](https://pre-commit.com/#config-additional_dependencies) | (optional) a list of dependencies that will be installed in the environment where this hook gets run. One useful application is to install plugins for hooks such as `eslint`.<br />（可选）在运行此钩子的环境中安装的依赖项列表。一个有用的应用是为诸如eslint之类的钩子安装插件。 |
| [`always_run`](https://pre-commit.com/#config-always_run)    | (optional) if `true`, this hook will run even if there are no matching files.<br />（可选）如果为真，则即使没有匹配的文件，此钩子也会运行。 |
| [`verbose`](https://pre-commit.com/#config-verbose)          | (optional) if `true`, forces the output of the hook to be printed even when the hook passes.<br />（可选）如果为真，则即使钩子通过，也会强制打印出钩子的输出。 |
| [`log_file`](https://pre-commit.com/#config-log_file)        | (optional) if present, the hook output will additionally be written to a file when the hook fails or [verbose](https://pre-commit.com/#config-verbose) is `true`.<br />（可选）如果存在，当钩子失败或 verbose 为 true 时，钩子输出还将被写入文件。 |
