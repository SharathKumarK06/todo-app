#!/usr/bin/env bash

docker run --name postgres \
  -e POSTGRES_USER=todo \
  -e POSTGRES_PASSWORD=pass \
  -e POSTGRES_DB=todo_db \
  -p 5432:5432 \
  -d postgres
