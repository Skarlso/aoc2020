#!/bin/zsh

for i in {1..25}; do
    mkdir -p day$i && mkdir -p day$i/part1 day$i/part2
    (
        cd day$i
        go mod init aoc2020/day$i
    )
done
