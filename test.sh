# !/bin/bash

set -x

curl -X PUT http://localhost:8080/profile?userEmail="dummy@gmail.com" -d'{"location":"dummy", "username":"dummy", "firstname":"dummy", "lastname":"dummy", "gender": "dummy", "dateOfBirth": "20/10/1990", "mobilenumber": "99xxxxxxx"}'
curl -X GET http://localhost:8080/profile?userEmail="dummy@gmail.com"
curl -X DELETE http://localhost:8080/profile?userEmail="dummy@gmail.com"
