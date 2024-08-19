# registry

注册中心实现接口，简单的一致性`hash`节点实现

**实现**

- [注册中心-redis](https://github.com/kordar/registry-redis)

**使用**

- `redis`[详见](https://github.com/kordar/registry-redis)

```go
redisnoderegistry := registry_redis.NewRedisNodeRegistry(client, &registry_redis.RedisNodeRegistryOptions{
    Prefix:  cast.ToString(item["prefix"]),
    Node:    cast.ToString(item["node"]),
    Timeout: timeout,
    Channel: cast.ToString(item["channel"]),
    Reload: func(value []string, channel string) {
        // ...
    },
    Heartbeat: heartbeat,
})

redisnoderegistry.Listener()
if err := redisnoderegistry.Register(); err == nil {
    logger.Infof("初始化registry redis成功")
} else {
    logger.Errorf("初始化registry redis异常，err=%v", err)
}
```

- 一致性`hash`使用

注，节点结构如下：
```json
{
  "node": "xxxx",
  "weight": "1"
}
```

```go
handle := registry.NewHashringRegistry(virtualSpots, item["node"])
handle.Load(value)
```