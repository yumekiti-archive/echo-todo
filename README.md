# echo-todo

## フレームワーク

https://echo.labstack.com/

## 依存関係

```
+----------------+
|   interface    |
+-------+--------+
        |
+-------v--------+
|    usecase     |
+-------+--------+
        |
+-------v--------+
|     domain     |
+-------^--------+
        |
+-------+--------+
| infrastructure |
+----------------+
```

## Example

```
curl -X GET "http://localhost:8080/task"

curl -X POST "http://localhost:8080/task" -H "Content-Type: application/json" -d '{"title": "ほげ", "body": "ふが"}'

curl -X GET "http://localhost:8080/task/1"

curl -X PUT "http://localhost:8080/task/1" -H "Content-Type: application/json" -d '{"title": "hoge", "body": "huga"}'

curl -X DELETE "http://localhost:8080/task/1"
```