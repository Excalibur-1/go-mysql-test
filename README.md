## Golang使用原生sql操作mysql数据库

* 数据库脚本：
```sql
create table user
(
    id               int(16) auto_increment
        primary key,
    email            varchar(255)        default ''                not null comment '邮箱',
    name             varchar(255)        default ''                not null comment '用户名',
    created_at       timestamp           default CURRENT_TIMESTAMP not null,
    updated_at       timestamp           default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
)
    comment '用户表' charset = utf8mb4;
```