#!/bin/bash
# @Author: wangsy
# @Date:   2018-01-15
# @Last Modified by:   wangsy


case $1 in 
	start)
		nohup ./gushici 2>&1 >> info.log 2>&1 /dev/null &
		echo "服务已启动..."
		sleep 1
	;;
	stop)
		killall gushici
		echo "服务已停止..."
		sleep 1
	;;
	restart)
		killall gushici
		sleep 1
		nohup ./gushici 2>&1 >> info.log 2>&1 /dev/null &
		echo "服务已重启..."
		sleep 1
	;;
	*) 
		echo "$0 {start|stop|restart}"
		exit 4
	;;
esac
