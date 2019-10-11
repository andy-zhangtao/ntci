# 快速入门

+ [NTCI配置](#ntci-config)
    1. [Buid Server配置](#build-config)
    2. [NTCI Server配置](#ntci)
+ [GitLab配置](#git-config)


<h3 id="ntci-config"> NTCI配置</h3>


<h4 id="build-config">Build Server(以k8s build Server为例)</h4>

* 创建参数文件
---
`k8s-build-server`默认使用`k8s.toml`作为运行参数文件，格式如下：

```toml
# 若为空, 默认使用80
port=8000

[k8s]
    [k8s.c1]
        endpoint="xx"
        token="xxx"
        config="xxx"

[language]
    go=[
        "vikings/go",
        "vikings/go:1.13.1"
    ]
```

* 快速启动
---

```shell script
docker run -it --rm --name ntci -p 8000:8000 -v k8s.toml:/k8s.toml vikings/k8s-build:latest
```

<h4 id="ntci">NTCI Server</h4>
* 创建参数文件
---
NTCI默认使用`ntci.toml`作为运行参数文件，文件格式如下:

```toml
# build mode
build-mode="single"
[access.gitlab]
    token="xxxxx"

[build]
    [build.single]
    addr="127.0.0.1:5000"
```

* 快速启动
---
```shell script
docker run -it --rm --name ntci -p 80:80 -v ntci.toml:/ntci.toml vikings/ntci:latest
```


<h3 id="git-config">Git配置</h3>

* GitLab
---
1. 创建一个具有`公共访问权限`的用户，并创建Access Token。
2. 在准备启用NTCI的工程中，添加此用户，并赋予`Reporter`角色。
3. 在`Integrations Settings`中，添加`Webhooks`并填写NTCI的访问URL，例如: `http://xxxxx/v1/gitlab/push`。 Trigger类型至少包括`Push Events`
4. 通过`Test`验证配置是否成功。 

