打印 所有线程堆栈
gdb 
thread apply all bt

You can attach to a running process with gdb -p PID.

Yes. You can do:

gdb program_name program_pid

A shortcut would be (assuming only one instance is running):

gdb program_name `pidof program_name`


删除 install 的库文件
xargs rm < install_manifest.txt

#include路径 问题
`gcc -print-prog-name=cc1plus` -v

#当前用户的include路径，一般自己安装的库放这里。比如PCL库，一般推荐放在这里
/usr/local/include
如何修改gcc、g++默认include路径？
方法一：通过命令行添加
g++ -I/home --I/tmp main.cpp  #/home的优先级高于/tmp
方法二：通过环境变量添加
#配置文件.bashrc是在/home/zss/，当前用户下，~就表示当前目录，你也可以通过指令 
locate  .bashrc
 
vim  ~/.bashrc
 
export CPLUS_INCLUDE_PATH = $CPLUS_INCLUDE_PATH:/Apollo
 
source ~/.bashrc
#最后需要让修改生效
 方法三：添加到/usr/local/include或者 /usr/include路径下
注意：

优先级排序：通过编译指定的include  > 环境变量 > /usr/local/include > /usr/include
