#!/bin/bash
go_program="./word_count.go"
dataset_path="$1"
go run "$go_program" "$dataset_path"
