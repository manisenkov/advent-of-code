#!/bin/bash

YEAR=$1
DAY=$(printf '%02d' $2)

cargo run --release --bin ${YEAR}_${DAY} < inputs/${YEAR}_${DAY}.txt
