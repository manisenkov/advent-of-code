#!/bin/bash

YEAR=$1
DAY=$(printf '%02d' $2)

poetry run pytest tests/year${YEAR}/test_day${DAY}.py
