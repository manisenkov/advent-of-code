#!/bin/bash

YEAR=$1
DAY=$(printf '%02d' $2)

cargo test --release --bin ${YEAR}_${DAY}
