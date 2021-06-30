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

data in channel : {2 62 dolorem cumque quo nihil inventore enim https://via.placeholder.com/600/65ad4f https://via.placeholder.com/150/65ad4f 62 966227}
data in channel : {1 40 est quas voluptates dignissimos sint praesentium nisi recusandae https://via.placeholder.com/600/3a0b95 https://via.placeholder.com/150/3a0b95 40 966227}
data in channel : {1 12 mollitia soluta ut rerum eos aliquam consequatur perspiciatis maiores https://via.placeholder.com/600/66b7d2 https://via.placeholder.com/150/66b7d2 12 966227}
data in channel : {4 200 perspiciatis est commodi iste nulla et eveniet voluptates eum https://via.placeholder.com/600/c3f384 https://via.placeholder.com/150/c3f384 200 966227}
data in channel : {2 52 eveniet pariatur quia nobis reiciendis laboriosam ea https://via.placeholder.com/600/121fa4 https://via.placeholder.com/150/121fa4 52 966227}
data in channel : {1 38 natus magnam iure rerum pariatur molestias dolore nisi https://via.placeholder.com/600/4f5b8d https://via.placeholder.com/150/4f5b8d 38 966227}
data in channel : {1 49 quasi quae est modi quis quam in impedit https://via.placeholder.com/600/2cd88b https://via.placeholder.com/150/2cd88b 49 966227}
data in channel : {1 21 ad et natus qui https://via.placeholder.com/600/5e12c6 https://via.placeholder.com/150/5e12c6 21 966227}
data in channel : {1 17 natus doloribus necessitatibus ipsa https://via.placeholder.com/600/9c184f https://via.placeholder.com/150/9c184f 17 966227}
data in channel : {1 18 laboriosam odit nam necessitatibus et illum dolores reiciendis https://via.placeholder.com/600/1fe46f https://via.placeholder.com/150/1fe46f 18 966227}
...
...
...
data in channel : {1 10 beatae et provident et ut vel https://via.placeholder.com/600/810b14 https://via.placeholder.com/150/810b14 10 966227}
data in channel : {1 23 harum velit vero totam https://via.placeholder.com/600/e924e6 https://via.placeholder.com/150/e924e6 23 966227}
data in channel : {4 177 ratione autem magni eveniet voluptas quia corporis https://via.placeholder.com/600/37942b https://via.placeholder.com/150/37942b 177 966227}
data in channel : {3 131 qui sed ex https://via.placeholder.com/600/af2618 https://via.placeholder.com/150/af2618 131 966227}
call api complete 
200
```

## Step to build and compile
- cd go_exam_3
- go build .
- run with ./go_exam_3 -id_range=200 -employee_id=966227

## Finally result
- should see '966227.xlsx' and snapshot image
- including column: GoRoutineId, EmployeeId, AlbumId, Id, Title, Url, ThumbnailUrl

## Issues
- change package from go_exam_3 to main for run this exam [DONE]