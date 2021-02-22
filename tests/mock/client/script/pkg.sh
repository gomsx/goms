#!/bin/bash

function getJsonValueByNode() {
	local json="$1"
	local key="$2"
	node -pe "JSON.stringify(JSON.parse(process.argv[1]).$key)" "$json" | awk '{print ($0 == "undefined" ? "null" : $0)}'
}

function getJsonValueByKey() {
	json="$1"
	key="$2"

	val="$(getJsonValueByNode "$json" "$key")"
	if [ -n "$val" ]; then
		echo "$val"
		return
	fi

	val=$(echo "$json" | jp "$key")
	if [ -n "$val" ]; then
		echo "$val"
		return
	fi

	echo "failed"
	return
}
