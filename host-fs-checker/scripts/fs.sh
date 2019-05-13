#!/bin/bash

mount=/fsmetrics
#outfile=/fsmetrics/fs.out
outfile=/fsmetricsOut/fs.out

while true 
do

# Note "df" command runs actually on the container - not on host
#df -k >> $outfile 2>> $outfile

# instead, use du -sk of the mounted volume
du -sk $mount > $outfile

sleep 10

done
