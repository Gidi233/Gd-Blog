## 日志规范
- 日志包统一使用 `github.com/Gidi233/Gd-Blog/internal/pkg/log`;
  - 日志包使用 [zap](https://github.com/uber-go/zap) 库,在 zap 库基础上进行了一层封装
  - 拥有 2 类 `zapLogger` 对象
    - 全局对象
      - 使用类似 `log.Infow()` 的方式进行调用
      - Init(opts *Options) 初始化全局对象
    - 局部对象
	  - 可以传入不同参数,创建一个自定义的 `*zapLogger` 对象
      - NewLogger(opts *Options) *zapLogger 创建一个自定义的 `*zapLogger`
  - 通过 Options 结构体对日志系统进行配置
    - 可以通过 NewOptions() 创建一个默认的 *Options 对象
    - 自定义配置文件存放于 yaml 之中,通过 `viper` 读取日志配置
- 使用结构化的日志打印格式：`log.Infow`, `log.Warnw`, `log.Errorw` 等; 例如：`log.Infow("Update post function called")`;
- 日志均以大写开头，结尾不跟 `.`，例如：`log.Infow("Update post function called")`;
- 使用过去时，例如：`Could not delete B` 而不是 `Cannot delete B`;
- 遵循日志级别规范：
  - Debug 级别的日志使用 `log.Debugw`;
  - Info 级别的日志使用 `log.Infow`;
  - Warning 级别的日志使用 `log.Warnw`;
  - Error 级别的日志使用 `log.Errorw`;
  - Panic 级别的日志使用 `log.Panicw`;
  - Fatal 级别的日志使用 `log.Fatalw`.
- 日志设置：
  - 开发测试环境：日志级别设置为 `debug`、日志格式可根据需要设置为 `console` / `json`、开启 caller；
  - 生产环境：日志级别设置为 `info`、日志格式设置为 `json`、开启 caller。（注意：上线初期，为了方便现网排障，日志级别可以设置为 `debug`）
- 在记录日志时，不要输出一些敏感信息，例如密码、密钥等。
- 如果在具有 `context.Context` 参数的函数/方法中，调用日志函数，建议使用 `log.C(ctx).Infow()` 进行日志记录。
