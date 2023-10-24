# 通用换电柜控制系统

目前集成：云动换电柜、凯信换电柜、西六楼换电柜以及小改动即可兼容其他使用铁塔换电协议的换电柜

### 第三方库

- [DTM是一款开源的分布式事务管理器](https://www.dtm.pub/guide/start.html)
- [kratos](https://go-kratos.dev/docs)
- [wire](https://zhuanlan.zhihu.com/p/399101012)
- [dig](https://darjun.github.io/2020/02/22/godailylib/dig/)
- [Loki](https://github.com/grafana/loki)
  - [Grafana Loki 开源日志聚合系统代替 ELK 或 EFK](https://wsgzao.github.io/post/loki/)
  - [Grafana Loki 简明教程](https://www.qikqiak.com/post/grafana-loki-usage/)
- [grpc-go-middleware](https://www.cnblogs.com/FireworksEasyCool/p/12750339.html)
- [GoKit CLI](https://github.com/GrantZheng/kit/blob/master/README_zh.md)
- [Go 教程：使用 GO-KIT 来构建微服务（上）](https://learnku.com/go/t/36923)
- [GO与PG实现缓存同步](https://pigsty.cc/zh/blog/2017/08/03/go%E4%B8%8Epg%E5%AE%9E%E7%8E%B0%E7%BC%93%E5%AD%98%E5%90%8C%E6%AD%A5/)
- [go| go并发实战: 搭配 influxdb + grafana 高性能实时日志监控系统](https://developer.aliyun.com/article/833106)

### 日志内容

- 时间
- 供应商
- 事件类型 (遥测 / 遥控)
- 事件描述

### 接口定义

- 管理端
    - 增删改查电柜
    - 控制电柜
- 骑手端

### kratos

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
```

### ent

```shell
go run -mod=mod entgo.io/ent/cmd/ent init --target ./internal/ent/schema CabinetBin
```

### postgres

- [Postgres Listen / Notify Real-time Notifications in Go](https://ds0nt.com/postgres-streaming-listen-notify-go)

```postgresql
SET TIMEZONE = 'Asia/Shanghai';
CREATE EXTENSION postgis;
CREATE EXTENSION pg_trgm;
CREATE EXTENSION btree_gin;
```

```postgresql
CREATE OR REPLACE FUNCTION notify_event() RETURNS TRIGGER AS
$$

DECLARE
    data         JSON;
    notification JSON;

BEGIN
    -- Convert the old or new row to JSON, based on the kind of action.
    -- Action = DELETE?             -> OLD row
    -- Action = INSERT or UPDATE?   -> NEW row
    IF (TG_OP = 'DELETE') THEN
        data = ROW_TO_JSON(OLD);
    ELSE
        data = ROW_TO_JSON(NEW);
    END IF;

    -- Contruct the notification as a JSON string.
    notification = JSON_BUILD_OBJECT(
            'table', TG_TABLE_NAME,
            'action', TG_OP,
            'data', data);

    -- Execute pg_notify(channel, notification)
    PERFORM pg_notify(TG_TABLE_NAME, notification::TEXT);

    -- Result is ignored since this is an AFTER trigger
    RETURN NULL;
END;

$$ LANGUAGE plpgsql;

DO
$$
    BEGIN
        IF NOT EXISTS(SELECT 1 FROM pg_trigger WHERE tgname = 'cabinet_notify_event') THEN
            CREATE TRIGGER cabinet_notify_event
                AFTER INSERT OR UPDATE OR DELETE
                ON cabinet
                FOR EACH ROW
            EXECUTE PROCEDURE notify_event();
        END IF;

        IF NOT EXISTS(SELECT 1 FROM pg_trigger WHERE tgname = 'bin_notify_event') THEN
            CREATE TRIGGER bin_notify_event
                AFTER INSERT OR UPDATE OR DELETE
                ON bin
                FOR EACH ROW
            EXECUTE PROCEDURE notify_event();
        END IF;
    END
$$;

CREATE OR REPLACE FUNCTION set_serial_id_seq() RETURNS TRIGGER AS
$$
BEGIN
    EXECUTE (FORMAT('SELECT setval(''%s_%s_seq'', (SELECT MAX(%s) from %s));',
                    TG_TABLE_NAME,
                    'id',
                    'id',
                    TG_TABLE_NAME));
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

DO
$$
    BEGIN
        CREATE OR REPLACE TRIGGER set_auto_id_seq
            AFTER INSERT OR UPDATE OR DELETE
            ON bin
            FOR EACH STATEMENT
        EXECUTE PROCEDURE set_serial_id_seq();

        CREATE OR REPLACE TRIGGER set_auto_id_seq
            AFTER INSERT OR UPDATE OR DELETE
            ON cabinet
            FOR EACH STATEMENT
        EXECUTE PROCEDURE set_serial_id_seq();
    END
$$;
```
