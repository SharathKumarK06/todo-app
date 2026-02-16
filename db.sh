#!/usr/bin/env bash

message="Usage: $(basename $0) [create|start|stop|db]"

[ "$#" != 1 ] && echo "$message" && exit 1

case "$1" in
  "create")
    docker run --name postgres \
      -e POSTGRES_USER=todo \
      -e POSTGRES_PASSWORD=pass \
      -e POSTGRES_DB=todo_db \
      -p 5432:5432 \
      -d postgres;
    ;;
  "start")
    docker start postgres
    ;;
  "stop")
    docker stop postgres
    ;;
  "db")
    docker exec -it postgres psql -U todo -d todo_db
    ;;
  *)
    echo "$message" && exit 1
    ;;
esac
