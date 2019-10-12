# 构建细节
> 内置在构建容器中的构建工具,负责拉取代码,并执行build

`Builder`运行所需参数是由`Build Server`创建Job时所提供。 

下面四个参数必须提供:

+ NTCI_BUILDER_JID 唯一Job ID.用于更新作业状态
+ NTCI_BUILDER_GIT Git仓库地址
+ NTCI_BUILDER_BRANCH 构建分支
+ NTCI_BUILDER_ADDR 构建服务地址

