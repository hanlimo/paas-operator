# Contributing 指南

**Note**: 本文档在持续编辑中

欢迎参与 paas-operator 项目！

作为这个项目的贡献者和维护者，为了打造一个实用和受欢迎的开源项目，我们欢迎任何人通过提出issues、feature requests、完善文档、提交各种 pull requests 等方式参与本项目。

## 开始编码前

### 开发环境

go version go1.10.3 linux/amd64

### 测试相关

### 日志规范

### 编码规范

和 kubernetes 看齐

### 文档规范

## 项目角色与职责

|   角色   |          职责           |        要求        |    指定     |
| :------: | :---------------------: | :----------------: | :---------: |
|  member  |      活跃的开发者       |  持续参与项目贡献  | 由owner指定 |
| reviewer | review 其他人提交的代码 |   成为member之后   | 由owner指定 |
|  owner   |       项目管理者        | 项目贡献代码量前三 |     N/A     |

所有 member 应该欢迎新加入的贡献者，积极帮助新贡献者熟悉项目贡献流程，协助新贡献者找到组织，加入微信群等。

member 列表维护在 [OWNERS](./OWNERS) 文件内

### Member

Member 是项目的活跃参与者，可以直接指定 issue 给自己，拥有直接往项目提交代码的权限（除了master外的分支）

### Reviewer

Reviewer 有资格决定一段代码是否能够合入master分支。

### Owner

Owner 可以给活跃的开发者 member 权限或者 reviewer 角色。

## 评论交流词规范

| 命令词汇                 | 谁可以用 |       详细说明        |
| :----------------------- | :------: | :-------------------: |
| /assign [@user1, @user2] |  member  | 指定一个issue由谁完成 |
| /cc [@user1]             |  任何人  |     请求reviewer      |
| /lgtm                    |  任何人  | 表示自己认可这份代码  |
| /approve                 | reviewer |    表示接受这个pr     |

## 提交第一个贡献

你可以在项目的 issue 列表中寻找自己感兴趣的议题，通过评论告诉 member 自己想要领取这个任务，如果 member 同意，会通过 /assign [@user] 的方式将这个任务指派给你。

同样你有好的点子，可以通过 issue 方式提交一个 feature 想法，大家讨论后决定要做，可以由你来实现。

## Github 工作流
