#!/bin/bash

# Create the input file with the following:
# find . -type d -links 2 -printf "%TY-%Tm-%TdT%TH:%TM:%.2TS %p\n" | sort -n > folder-modified-dates.txt

while read -r date filepath;
do
    echo "$date $filepath";
    if [[ $filepath == *"/."* ]]; then
      echo "dot folder, skipping";
    else
      git add "$filepath";
      GIT_COMMITTER_DATE="$date" GIT_AUTHOR_DATE="$date" git commit -m "added $(echo "$filepath" | cut -c3-)";
    fi;
done < folder-modified-dates.txt;
