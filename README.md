<<<<<<< HEAD
## server项目结构

```shell
├── api
│   └── v1
├── config
├── core
├── docs
├── global
├── initialize
│   └── internal
├── middleware
├── model
│   ├── request
│   └── response
├── packfile
├── resource
│   ├── excel
│   ├── page
│   └── template
├── router
├── service
├── source
└── utils
    ├── timer
    └── upload
```

| 文件夹       | 说明                    | 描述                        |
| ------------ | ----------------------- | --------------------------- |
| `api`        | api层                   | api层 |
| `--v1`       | v1版本接口              | v1版本接口                  |
| `config`     | 配置包                  | config.yaml对应的配置结构体 |
| `core`       | 核心文件                | 核心组件(zap, viper, server)的初始化 |
| `docs`       | swagger文档目录         | swagger文档目录 |
| `global`     | 全局对象                | 全局对象 |
| `initialize` | 初始化 | router,redis,gorm,validator, timer的初始化 |
| `--internal` | 初始化内部函数 | gorm 的 longger 自定义,在此文件夹的函数只能由 `initialize` 层进行调用 |
| `middleware` | 中间件层 | 用于存放 `gin` 中间件代码 |
| `model`      | 模型层                  | 模型对应数据表              |
| `--request`  | 入参结构体              | 接收前端发送到后端的数据。  |
| `--response` | 出参结构体              | 返回给前端的数据结构体      |
| `packfile`   | 静态文件打包            | 静态文件打包 |
| `resource`   | 静态资源文件夹          | 负责存放静态文件                |
| `--excel` | excel导入导出默认路径 | excel导入导出默认路径 |
| `--page` | 表单生成器 | 表单生成器 打包后的dist |
| `--template` | 模板 | 模板文件夹,存放的是代码生成器的模板 |
| `router`     | 路由层                  | 路由层 |
| `service`    | service层               | 存放业务逻辑问题 |
| `source` | source层 | 存放初始化数据的函数 |
| `utils`      | 工具包                  | 工具函数封装            |
| `--timer` | timer | 定时器接口封装 |
| `--upload`      | oss                  | oss接口封装        |

=======
### 3 分钟了解如何进入开发

欢迎使用 Codeup，通过阅读以下内容，你可以快速熟悉 Codeup ，并立即开始今天的工作。

### 提交**文件**

首先，你需要了解在 Codeup 中如何提交代码文件，跟着文档「[__提交第一行代码__](https://thoughts.teambition.com/sharespace/5d88b152037db60015203fd3/docs/5dc4f6786b81620014ef7574)」一起操作试试看吧。

### 开启扫描

开发过程中，为了更好的管理你的代码资产，Codeup 内置了「[__代码规约扫描__](https://thoughts.teambition.com/sharespace/5d88b152037db60015203fd3/docs/5dc4f68b6b81620014ef7588)」和「[__敏感信息检测__](https://thoughts.teambition.com/sharespace/5d88b152037db60015203fd3/docs/5dc4f6886b81620014ef7587)」服务，你可以在代码库设置-集成与服务中一键开启，开启后提交或合并请求的变更将自动触发扫描，并及时提供结果反馈。

![](https://img.alicdn.com/tfs/TB1nRDatoz1gK0jSZLeXXb9kVXa-1122-380.png "")

![](https://img.alicdn.com/tfs/TB1PrPatXY7gK0jSZKzXXaikpXa-1122-709.png "")

### 代码评审

功能开发完毕后，通常你需要发起「[__代码合并和评审__](https://thoughts.teambition.com/sharespace/5d88b152037db60015203fd3/docs/5dc4f6876b81620014ef7585)」，Codeup 支持多人协作的代码评审服务，你可以通过「[__保护分支__](https://thoughts.teambition.com/sharespace/5d88b152037db60015203fd3/docs/5dc4f68e6b81620014ef758c)」策略及「[__合并请求设置__](https://thoughts.teambition.com/sharespace/5d88b152037db60015203fd3/docs/5dc4f68f6b81620014ef758d)」对合并过程进行流程化管控，同时提供 WebIDE 在线代码评审及冲突解决能力，让你的评审过程更加流畅。

![](https://img.alicdn.com/tfs/TB1XHrctkP2gK0jSZPxXXacQpXa-1432-887.png "")

![](https://img.alicdn.com/tfs/TB1V3fctoY1gK0jSZFMXXaWcVXa-1432-600.png "")

### 编写文档

项目推进过程中，你的经验和感悟可以直接记录到 Codeup 代码库的「[__文档__](https://thoughts.teambition.com/sharespace/5d88b152037db60015203fd3/docs/5e13107eedac6e001bd84889)」内，让智慧可视化。

![](https://img.alicdn.com/tfs/TB1BN2ateT2gK0jSZFvXXXnFXXa-1432-700.png "")

### 成员协作

是时候邀请成员一起编写卓越的代码工程了，请点击右上角「成员」邀请你的小伙伴开始协作吧！

### 更多

Git 使用教学、高级功能指引等更多说明，参见[__Codeup帮助文档__](https://thoughts.teambition.com/sharespace/5d88b152037db60015203fd3/docs/5dc4f6756b81620014ef7571)。
>>>>>>> c3b2a07f7173a368a9511539e7642f75ae98d1e7
