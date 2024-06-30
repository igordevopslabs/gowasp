#! /bin/bash

for i in {1..100}; do curl -X GET http://localhost:3000/ping; done
