# GoLang App

## Curl Request
* POST
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"task1", "context":"context1", "tag":"tag1"}' localhost:8001/api/v1/tasks
```
* GET ALL
```
curl -X GET -H "Content-Type: application/json" localhost:8001/api/v1/tasks
```
* GET BY ID
```
curl -X GET -H "Content-Type: application/json" localhost:8001/api/v1/tasks/<ID>
```
* Update
```
curl -X PUT -H "Content-Type: application/json" -d '{"name":"task1_fixed", "context":"context1_fixed", "tag":"tag1_fixed"}' localhost:8001/api/v1/tasks/1
```
* DELETE
```
curl -X DELETE localhost:8001/api/v1/tasks/4
```
