# chatforum
online chat

## 在线论坛


-   main.go：应用入口文件
-   config.json：全局配置文件
-   handlers：用于存放处理器代码（可类比为 MVC 模式中的控制器目录）
-   logs：用于存放日志文件
-   models：用于存放与数据库交互的模型类
-   public：用于存放前端资源文件，比如图片、CSS、JavaScript 等
-   routes：用于存放路由文件和路由器实现代码
-   views：用于存放视图模板文件


在`chatforum`数据库中创建相应数据表，对应SQL语句如下：
```mysql
create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null
);
    
create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null
);
    
create table threads (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  topic      text,
  user_id    integer references users(id),
  created_at timestamp not null
);
    
create table posts (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  thread_id  integer references threads(id),
  created_at timestamp not null
);
```


## 访问的整体流程
用户请求的处理流程如下：

1. 客户端发送请求；
2. 服务端路由器（multiplexer）将请求分发给指定处理器（handler）；
3. 处理器处理请求，完成对应的业务逻辑；
3. 处理器调用模板引擎生成 HTML 并将响应返回给客户端。

