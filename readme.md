> go-micro 相关测试



---

### 前置物料

因为使用 模块化. 在 *go.mod* 中指向了本地包

```json
replace github.com/micro/go-micro => ../github.com/micro/go-micro
```

所有, 要正常跑起来有两种姿势

1. 简单版: 直接将这个 replace 注释掉.

2. 复制版: 把 GitHub 上的包 完整 下下来. 可以通过子模块的方式添加进来. (GitHub地址: `git@github.com:micro/go-micro.git`)

    ![](http://yxbl.itengshe.com/20191012163315-1.png)