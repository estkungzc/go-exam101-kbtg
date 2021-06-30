# go_exam_3

## Flag
- Required id_range; 0 < id_range < 250 (default: 0)
- Required employee_id (default: "")

## Step to run go code (Normal)
- cd go_exam_3
- go run .

```shell
# example result

$ go run . -id_range=200 -employee_id=966227

data in channel : {1 12 mollitia soluta ut rerum eos aliquam consequatur perspiciatis maiores https://via.placeholder.com/600/66b7d2 https://via.placeholder.com/150/66b7d2 12 966227 <nil>}
data in channel : {2 94 et quisquam aspernatur https://via.placeholder.com/600/cc12f5 https://via.placeholder.com/150/cc12f5 94 966227 <nil>}
data in channel : {1 25 facere non quis fuga fugit vitae https://via.placeholder.com/600/5e3a73 https://via.placeholder.com/150/5e3a73 25 966227 <nil>}
data in channel : {1 16 iusto sunt nobis quasi veritatis quas expedita voluptatum deserunt https://via.placeholder.com/600/fdf73e https://via.placeholder.com/150/fdf73e 16 966227 <nil>}
...
...
...
data in channel : {4 172 deserunt commodi et aut et molestiae debitis et sed https://via.placeholder.com/600/d611bd https://via.placeholder.com/150/d611bd 172 966227 <nil>}
data in channel : {3 143 totam temporibus eaque est eum et perspiciatis ullam https://via.placeholder.com/600/79439b https://via.placeholder.com/150/79439b 143 966227 <nil>}
data in channel : {4 167 et fugit et eius quod provident https://via.placeholder.com/600/3b4a81 https://via.placeholder.com/150/3b4a81 167 966227 <nil>}
call api complete 
200
```

## Step to build and compile
- cd go_exam_3
- go build .
- run with ./go_exam_3 -id_range=200 -employee_id=966227

## Finally result
- should see '966227.xlsx' and snapshot image
- including column: GoRoutineId, EmployeeId, AlbumId, Id, Title, Url, ThumbnailUrl, Error
- Maybe write back error to excel if them found

## Issues
- change package from go_exam_3 to main for run this exam [DONE]