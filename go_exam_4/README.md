# go_exam_4

## Environment
1. including local(default), dev
2. set ENV VARIABLE `STATE` to `dev` for change environment

## How to run
1. Stay in the root project (current directory is go-exam101-kbtg)
2. Run with `go run go_exam_4/main.go` to start Echo
4. Test api on localhost:8888 (employee/byId, employee/byFirstName)
```shell
# example result call api


POST http://localhost:8888/employee/byFirstName

HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 30 Jun 2021 08:45:26 GMT
Content-Length: 192

{
  "success": true,
  "code": 200,
  "message": "ok",
  "data": [
    {
      "first_name": "Gopher",
      "last_name": "Conference",
      "emp_id": "0000000001"
    },
    {
      "first_name": "Gopher",
      "last_name": "Cons2020",
      "emp_id": "9999999999"
    }
  ]
}


Response code: 200 (OK); Time: 267ms; Content length: 192 bytes
```

```shell
# example log from Echo to show env from yaml file
# test with flag

{"time":"2021-06-30T15:40:17.55805+07:00","level":"INFO","prefix":"-","file":"config.go","line":"43","message":"config file using : config.dev"}
{"time":"2021-06-30T15:40:18.116796+07:00","level":"INFO","prefix":"-","file":"config.go","line":"73","message":"all config loaded : &internal.Configs{vn:(*viper.Viper)(0xc000502300), ConfigPath:\"./go_exam_4/configs\", State:\"dev\", Validator:(*validator.Validate)(nil), TimeZone:\"Asia/Bangkok\", MongoDB:internal.MongoDB{MongoUri:\"mongodb+srv://%s:%s@cluster0.5m0mj.mongodb.net/%s?retryWrites=true&w=majority\", Timeout:5000000000, Username:\"teerapat\", Password:\"HDrLv7gRVMJ8BJaX\", Database:\"sample_employee\", Client:(*mongo.Client)(0xc0000b4e00)}, BangkokLocation:(*time.Location)(0xc0001d4310)}"}
```