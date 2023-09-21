#!/bin/bash

root_directory="../dataset"

total_count=0

count_words() {
	local directory="$1"
	temp_file="/tmp/count_$(echo "$directory" | tr / _)"
	count=$(python3 word_count.py "$directory")
	echo "$count" > "$temp_file"
}
directories=($(find "$root_directory" -type d))
num_processes=4
bg_processes=()

for directory in "${directories[@]}"; do
    if [ "$directory" != "$root_directory" ]; then
        count_words "$directory" &
        bg_processes+=("$!")
    fi
done

for pid in "${bg_processes[@]}"; do
    wait "$pid"
done

for directory in "${directories[@]}"; do
	if [ "$directory" != "$root_directory" ]; then
		temp_file="/tmp/count_$(echo "$directory" | tr / _)"
		count=$(cat "$temp_file" 2>/dev/null)
		if [ -n "$count" ]; then
			total_count=$((total_count + count))
			rm "$temp_file"
		fi
	fi
done
echo "The total directory has $total_count words"
