#!/bin/sh
# compile and run in the background
# log saved to torpedo.log

go build
./torpedo > torpedo.log 2>&1 &

pidof ./torpedo > torpedo.pid
