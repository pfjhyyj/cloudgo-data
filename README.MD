# 服务计算cloudgo-data

学号：15331366

使用orm使大部分的dao操作都实现了自动化

# post测试

```
$curl -v  http://localhost:8080/service/userinfo -d "username=123&departname=456"
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /service/userinfo HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
> Content-Length: 27
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 27 out of 27 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Wed, 29 Nov 2017 11:35:47 GMT
< Content-Length: 110
<
{
  "UID": 1,
  "UserName": "123",
  "DepartName": "456",
  "CreateAt": "2017-11-29T19:35:47.8703154+08:00"
}
* Connection #0 to host localhost left intact
```

# get测试

```
$ curl -v  http://localhost:8080/service/userinfo?userid=1
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /service/userinfo?userid=1 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 400 Bad Request
< Content-Type: application/json; charset=UTF-8
< Date: Wed, 29 Nov 2017 11:50:23 GMT
< Content-Length: 102
<
{
  "UID": 1,
  "UserName": "123",
  "DepartName": "456",
  "CreateAt": "2017-11-30T03:35:47+08:00"
}
```


# ab测试（orm）

```
$ ab -n 1000 -c 100 http://localhost:8080/service/userinfo?userid=
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        233 bytes

Concurrency Level:      100
Time taken for tests:   0.973 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      357000 bytes
HTML transferred:       233000 bytes
Requests per second:    1027.37 [#/sec] (mean)
Time per request:       97.336 [ms] (mean)
Time per request:       0.973 [ms] (mean, across all concurrent requests)
Transfer rate:          358.17 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       2
Processing:     7   93  26.8     92     171
Waiting:        5   73  24.0     74     146
Total:          7   93  26.8     92     171

Percentage of the requests served within a certain time (ms)
  50%     92
  66%    102
  75%    108
  80%    117
  90%    126
  95%    142
  98%    155
  99%    162
 100%    171 (longest request)
 ```

 # ab测试（database/sql）

 ```
$ab -n 1000 -c 100 http://localhost:8080/service/userinfo?userid=
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        223 bytes

Concurrency Level:      100
Time taken for tests:   1.245 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      347000 bytes
HTML transferred:       223000 bytes
Requests per second:    803.01 [#/sec] (mean)
Time per request:       124.531 [ms] (mean)
Time per request:       1.245 [ms] (mean, across all concurrent requests)
Transfer rate:          272.11 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       3
Processing:    38  120  54.8    107     296
Waiting:       14  107  60.0     96     292
Total:         39  120  54.8    107     296

Percentage of the requests served within a certain time (ms)
  50%    107
  66%    123
  75%    135
  80%    148
  90%    190
  95%    267
  98%    285
  99%    288
 100%    296 (longest request)
 ```

可以看出，高并发时orm效率更高