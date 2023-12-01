#!/bin/bash

YEAR=$1
DAY=$(printf '%02d' $2)

poetry run python -m adventofcode.year${YEAR}.day${DAY} < inputs/year${YEAR}/day${DAY}.txt
