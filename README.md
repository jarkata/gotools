go common tools

## date 格式化及解析

## snowflake id 生成

```golang
	nodeId := int64(1)
	workId := int64(2)
	id := tid.New(nodeId, workId)
	slog.Info("", "ID", id)

```

## 工具类
