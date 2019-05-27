#!/usr/bin/env bash

clear;
go build -o ./client.bin ./client ;
host=${HOST:-localhost};
port=${PORT:-8000};
./client.bin -host $host -port $port;

