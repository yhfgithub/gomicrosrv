cd ./proto
protoc --micro_out=./ --go_out=./ *.proto
cd ..