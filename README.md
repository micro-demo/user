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