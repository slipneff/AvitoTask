#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
user="$2"
password="$3"
dbname="$4"
port="$5"
shift 5
cmd="$@"

until PGPASSWORD=$password psql -h "$host" -U "$user" -d "$dbname" -p "$port" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 5
done

>&2 echo "Postgres is up - executing command"
exec $cmd
