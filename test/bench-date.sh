#!/bin/sh

sed '/^\s*#/d' test/samples/tz.conf |
while IFS= read -r l
do
\tTZ=$l date
done
