#!/bin/sh

for i in `seq 1 10`
do
    curl -d "hello $i" 'http://localhost:4151/pub?topic=test_topic'
done
