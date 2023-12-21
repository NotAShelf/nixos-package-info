#!/usr/bin/env bash

# Help function to display usage instructions
help_function() {
	echo "Usage: $0 <json_file>"
	echo "  <json_file>       JSON file to update"
	exit 1
}

# Fetch the latest commit from the GitHub API
fetch_latest_commit() {
	curl_output=$(curl -s https://api.github.com/repos/NixOS/nixpkgs/commits/master)
	latest_commit=$(jq -r .sha <<<"$curl_output")
}

# Update the git_ref in targets.json with the latest commit
update_targets_json() {
	jq --arg latest_commit "$latest_commit" '.[0].git_ref = $latest_commit' <"$1" >temp.json
	mv temp.json "$1"
}

# Check for arguments and display usage if help command or no arguments are given
if [[ $# -eq 0 || $1 == "-h" || $1 == "--help" ]]; then
	help_function
fi

# Check if the provided argument is a valid JSON file
if ! [[ -f "$1" && "${1##*.}" == "json" ]]; then
	echo "Error: Please provide a valid JSON file."
	help_function
fi

# Main script logic
fetch_latest_commit
update_targets_json "$1"
