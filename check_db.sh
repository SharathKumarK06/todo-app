#!/usr/bin/env bash

docker exec -it postgres psql -U todo -d todo_db
