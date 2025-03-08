go build -o go-https-server src/*
nohup ./go-https-server > logs.txt 2>&1 &
tail -f logs.txt