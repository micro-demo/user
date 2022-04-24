# User Service

This is the User service

Generated with

```go
micro server //保持micro server是启动状态的
```

```
micro new user
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```

> 如果有问题micro status 状态是error，大概率是micro server有问题。比如你切换wifi了或者挂梯子了

小写访问/user/userinfo
路由参数/user/userinfo/1
现在是post get都能访问我是否能指定http 请求方式？