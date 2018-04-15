# goecho
simple tcp socket listener echos back to all connected clients

author: mathias.dietrich@devops.ooo

Listents on TCP  port 2000 



compile
=======

go build echo.go
./echo


Test
====

nc localhost 2000
telnet localhost 2000


Check Latency
=============

nmap -p 2000 remotehost
