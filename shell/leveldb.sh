leveldb 编译动态库
sed -i "s/add_library(leveldb \"\")/add_library(leveldb \"SHARED\")/g" CMakeLists.txt