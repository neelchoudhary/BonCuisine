#!/bin/bash -e

echo $APP_ENV
APP_ENV=${APP_ENV:-local}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."

#echo "[`date`] Running DB migrations..."
#migrate -database "${APP_DSN}" -path ./migrations up

echo "[`date`] Starting server..."

go run cmd/server.go -env ${APP_ENV}