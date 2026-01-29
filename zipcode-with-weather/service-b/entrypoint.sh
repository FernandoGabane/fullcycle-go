#!/bin/sh

echo "Aguardando otel-collector:4317..."
until nc -z otel-collector 4317; do
  sleep 1
done

echo "Iniciando service-b..."
exec ./main
