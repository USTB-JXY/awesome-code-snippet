你可以使用以下的shell脚本来查看k-server程序的pid:
#!/bin/bash
ps -ef | grep k-server | grep -v grep | awk '{print $2}'
复制
这个脚本会列出所有的k-server进程并且输出它们的pid。

你可以使用以下的shell脚本来在后台模式运行k-server程序，并在程序崩溃后自动重启该程序:
#!/bin/bash
while true; do
  ./k-server &
  pid=$!
  wait $pid
done


你可以使用以下的shell脚本来在后台模式运行Bash脚本，并将输出重定向到空:
#!/bin/bash
nohup bash test.sh >/dev/null 2>&1 &
复制
这个脚本会在后台模式运行Bash脚本，并将输出重定向到空。

消除Windows换行
sed -i -e 's/\r$//' restart.sh 

你可以使用以下的命令来打开1300端口:

sudo iptables -A INPUT -p tcp --dport 1300 -j ACCEPT


查看系统运行时间
cat /proc/uptime| awk -F. '{run_days=$1 / 86400;run_hour=($1 % 86400)/3600;run_minute=($1 % 3600)/60;run_second=$1 % 60;printf("系统已运行：%d天%d时%d分%d秒",run_days,run_hour,run_minute,run_second)}'

date -d "$(awk -F. '{print $1}' /proc/uptime) second ago" +"%Y-%m-%d %H:%M:%S" 

统计当前目录文件大小
du -h -x --max-depth=1




