#!/bin/bash

YEAR=$1
DAY=$(printf '%02d' $2)

cargo run --bin ${YEAR}_${DAY} < inputs/${YEAR}_${DAY}.txt
