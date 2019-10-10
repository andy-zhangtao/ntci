# 快速入门

### NTCI配置

* 创建参数文件

NTCI默认使用`ntci.toml`作为运行参数文件，文件格式如下:

```toml

```

### Git配置

* GitLab
---
1. 创建一个具有`公共访问权限`的用户，并创建Access Token。
2. 在准备启用NTCI的工程中，添加此用户，并赋予`Reporter`角色。
3. 在`Integrations Settings`中，添加`Webhooks`并填写NTCI的访问URL，例如: `http://xxxxx/v1/gitlab/push`。 Trigger类型至少包括`Push Events`
4. 通过`Test`验证配置是否成功。 

