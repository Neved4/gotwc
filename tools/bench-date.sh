#!/bin/sh

sed '/^\s*#/d' tools/samples/tz.conf |
while IFS= read -r l
do
    TZ=$l date
done
