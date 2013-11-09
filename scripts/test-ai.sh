#!/bin/sh

# build and run
go build
./torpedo -config scripts/dev.config > test.log 2>&1 &
pidof ./torpedo > test.pid

# make sure the server started.
sleep 1

# add 2 players
curl --data "username=t1&is_robot=on" localhost:8080/join
curl --data "username=t2&is_robot=on" localhost:8080/join

# just to be sure that the game runs, wait a bit
sleep 1

# stop the server
cat test.pid | xargs kill
# clean up
rm test.pid

# show me the money
tail -20 test.log
