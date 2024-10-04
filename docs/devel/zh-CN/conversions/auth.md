# 认证与授权

## 认证 - Authentication - Authn
- 采用`基础认证` + `令牌认证`
  - 基础认证: 用户名+密码
  - 令牌认证: JWT
![JWT认证流程](../../../../internal/resource/JWT.png)

### JWT
- Json Web Token
- 由三部分组成: Header, Payload, Signature
- 一个示例:
```text
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ
```
- Header
  - 由两部分组成:
    - Token 的类型
    - Token 采用的加密算法
  - 例如:
```json
{
  "typ": "JWT",
  "alg": "HS256"
}
```
- Payload
  - 由多个键值对组成
    - iss: 签发者
    - sub: 主题
    - aud: 受众
    - exp: 过期时间
    - nbf: 生效时间
    - iat: 签发时间
    - jti: 唯一身份标识
    - 自定义键值对
  - 例如:
```json
{
  "id": 2,
  "username": "daz",
  "nbf": 1527931805,
  "iat": 1527931805
}
```
- Signature
  - 生成方式:
    - 用 Base64 对 Header 和 Payload 进行编码
    - 用 Secret 对编码后的 Header 和 Payload 进行加密, 生成 Signature
  - Secret 相当于一个私钥, 只有服务端知道
    - Gd-Blog 将其存放于 config/GdBlog.yaml 中

### 认证流程
- [token](../../../../pkg/token/token.go) 签发以及解析密钥
- [Authn](../../../../internal/pkg/middleware/authn.go) 中间件进行认证
- 基于以上实现 `POST /login` 以及 `PUT /v1/users/:name/change-password` 接口
  - POST 用于创建
  - PUT 用于更新,是一个幂等操作

# casbin授权 - Authorization - Authz
- 使用 ACL (Access Control List) 模型进行授权
- 
## 元模型
- PERM 元模型
- policy, effect, request, matcher
  - sub: subject, 访问实体
  - obj: object, 被访问实体
  - act: action, 访问行为
  - eft: effect, 访问结果,一般为空,默认指定为 allow 或 deny

### Policy
- 策略
- p = {sub, obj, act, eft}
- 一般存储与数据库中,因为会有很多
```text
[policy_definition]
p = sub, obj, act, eft
```


### Matchers
- 匹配规则
```text
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
```
- r 是请求, p 是策略
- 会把 r 和 p 按照上述描述进行匹配
  - 从而返回匹配结果(eft),如果不定义,则返回allow,否则返回定义值


### Effect
- 影响
- 决定我们受否放行
- cabin 支持的 policy effect

| Policy effect                                               | 意义                   |
| ----------------------------------------------------------- | ---------------------- |
| some(where (p.eft == allow))                                | allow-override         |
| !some(p.eft == deny)                                        | deny-override          |
| some(where (p.eft == allow) && !some(where (p.eft == deny)) | allow-and-deny         |
| priority(p.eft)\|\|deny                                     | priority               |
| subjectPriority(p.eft)                                      | priority based on role |

### Resource
- 请求
- r = {sub, obj, act}


## 角色域
- role_definition
  - g = _, _ 表示以角色为基础
  - g = _, _, _ 表示以域为基础(多商户模式)


## Gd-Blog 对 casbin 的使用
- 使用 RBAC 模型
  - Role Based Access Control
- 对资源操作进行授权
  - `用户`只可以访问自己账户下的 用户/博客 等资源
  - `管理员`可以访问所有资源
  - 也就是对 API 路径进行授权
- 授权策略:   

| A   | B    | C             | D                        |
| --- | ---- | ------------- | ------------------------ |
| p   | root | /v1/users*    | (GET)(POST)(PUT)(DELETE) |
| p   | bob  | /v1/users/bob | (GET)(POST)(PUT)(DELETE) |

- 因为要对每一个 HTTP 进行授权, 所以将授权功能封装为中间件