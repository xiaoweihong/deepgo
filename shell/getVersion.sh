#!/usr/bin/env bash

if [ $# -ne 1 ];then
    echo "please input release name"
    exit
fi

release=$1
#curl -s 192.168.2.136/software/release/deepvideo_release/$release/|awk -F "${release}-" 'NR>5 {print $2}'|awk '{if(!NF ){next}}1'|awk -F '\"' '{print $1}'
#curl -s 192.168.2.136/software/release/deepvideo_release/crusader/|awk -F "crusader-" 'NR>5 {print $2}'|awk '{if(!NF ){next}}1'|awk -F '\"' '{print $1}'
curl -s 192.168.2.136/software/release/deepvideo_release/$release/|awk  'NF>4{print $2,$3}'|awk -F "[=\" ]+" '{print $2,$4}'|grep $release