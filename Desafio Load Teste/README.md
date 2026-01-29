# Go Load Tester CLI

## Build
docker build -t go-load-tester .

## Run
docker run go-load-tester --url=http://google.com --requests=100 --concurrency=10
