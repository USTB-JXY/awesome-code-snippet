sed -i "s/127.0.0.1/192.168.145.61/g" ./dj.config

leveldb 编译动态库
sed -i "s/add_library(leveldb \"\")/add_library(leveldb \"SHARED\")/g" CMakeLists.txt