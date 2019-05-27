#!/usr/bin/env bash

clear;
host=${HOST:-localhost}
port=${PORT:-8000}
go build -o ./server.bin ./server;
./server.bin -host $host -port $port;

