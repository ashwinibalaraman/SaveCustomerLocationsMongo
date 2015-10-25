# cmpe-273-assignment2

Output:


/*************POST*************/

curl -vH "Content-Type: application/json" -X POST -d '{"name": "Abc Xyz",                                                                                
    "address": "123 Main St",
    "city": "San Francisco",
    "state": "CA",
    "zip": "94113"
}' http://127.0.0.1:8080/locations
* About to connect() to 127.0.0.1 port 8080 (#0)
*   Trying 127.0.0.1...
* connected
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> POST /locations HTTP/1.1
> User-Agent: curl/7.27.0
> Host: 127.0.0.1:8080
> Accept: */*
> Content-Type: application/json
> Content-Length: 123
> 
* upload completely sent off: 123 out of 123 bytes
< HTTP/1.1 200 OK
< Date: Sun, 25 Oct 2015 05:34:06 GMT
< Content-Length: 184
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 127.0.0.1 left intact
OK{"Id":"562c69ce91643f04228518bc","Name":"Abc Xyz","Address":"123 Main St","City":"San Francisco","State":"CA","Zip":"94113","Coord":{"Lat":"3.77917618E+01","Lng":"-1.223943405E+02"}}* Closing connection #0

/************GET*************/


curl -vH "Content-Type: application/json" -X GET http://127.0.0.1:8080/locations/562c69ce91643f04228518bc
* About to connect() to 127.0.0.1 port 8080 (#0)
*   Trying 127.0.0.1...
* connected
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /locations/562c69ce91643f04228518bc HTTP/1.1
> User-Agent: curl/7.27.0
> Host: 127.0.0.1:8080
> Accept: */*
> Content-Type: application/json
> 
< HTTP/1.1 200 OK
< Date: Sun, 25 Oct 2015 05:35:27 GMT
< Content-Length: 182
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 127.0.0.1 left intact
{"Id":"562c69ce91643f04228518bc","Name":"Abc Xyz","Address":"123 Main St","City":"San Francisco","State":"CA","Zip":"94113","Coord":{"Lat":"3.77917618E+01","Lng":"-1.223943405E+02"}}* Closing connection #0

/*************PUT************/


curl -vH "Content-Type: application/json" -H 'Accept: application/json' -X PUT -d '{
    "address": "1600AmphitheatreParkway",
    "city": "MountainView",
    "state": "CA",
    "zip": "94043"
}' http://127.0.0.1:8080/locations/562c69ce91643f04228518bc
* About to connect() to 127.0.0.1 port 8080 (#0)
*   Trying 127.0.0.1...
* connected
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> PUT /locations/562c69ce91643f04228518bc HTTP/1.1
> User-Agent: curl/7.27.0
> Host: 127.0.0.1:8080
> Content-Type: application/json
> Accept: application/json
> Content-Length: 111
> 
* upload completely sent off: 111 out of 111 bytes
< HTTP/1.1 201 Created
< Date: Sun, 25 Oct 2015 05:40:36 GMT
< Content-Length: 200
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 127.0.0.1 left intact
Created{"Id":"562c69ce91643f04228518bc","Name":"Abc Xyz","Address":"1600AmphitheatreParkway","City":"MountainView","State":"CA","Zip":"94043","Coord":{"Lat":"3.77917618E+01","Lng":"-1.223943405E+02"}}* Closing connection #0



/************** DELETE ***************/


curl -vX DELETE http://127.0.0.1:8080/locations/562c69ce91643f04228518bc
* About to connect() to 127.0.0.1 port 8080 (#0)
*   Trying 127.0.0.1...
* connected
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> DELETE /locations/562c69ce91643f04228518bc HTTP/1.1
> User-Agent: curl/7.27.0
> Host: 127.0.0.1:8080
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Sun, 25 Oct 2015 05:41:44 GMT
< Content-Length: 2
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 127.0.0.1 left intact
OK* Closing connection #0



/******* GET after DELETE **********/

curl -vH "Content-Type: application/json" -X GET http://127.0.0.1:8080/locations/562c69ce91643f04228518bc
* About to connect() to 127.0.0.1 port 8080 (#0)
*   Trying 127.0.0.1...
* connected
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /locations/562c69ce91643f04228518bc HTTP/1.1
> User-Agent: curl/7.27.0
> Host: 127.0.0.1:8080
> Accept: */*
> Content-Type: application/json
> 
< HTTP/1.1 404 Not Found
< Date: Sun, 25 Oct 2015 05:42:42 GMT
< Content-Length: 9
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 127.0.0.1 left intact
Not Found* Closing connection #0
















