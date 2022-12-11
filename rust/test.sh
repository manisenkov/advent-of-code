#!/bin/bash

YEAR=$1
DAY=$(printf '%02d' $2)

cargo test --bin ${YEAR}_${DAY}
