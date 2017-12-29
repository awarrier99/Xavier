#!/usr/bin/env bash
X="export XAVIER=\$GOPATH/src/github.com/awarrier99/Xavier"
G="export GOOGLE_APPLICATION_CREDENTIALS=\$XAVIER/keys/Xavier-be41d74a5f4f.json"
A="export AWS_ACCESS_KEY_ID=\"AKIAIVX5SSNXQH5MPDKA\""
S="export AWS_SECRET_KEY=\"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\""

LINES=("$X" "$G" "$A" "$S")


for ((i = 0; i < ${#LINES[@]}; i++)); do
	TEMP="${LINES[$i]}"
	if grep -Fxq "$TEMP" ~/.bash_profile; then
		:
	else
		echo "$TEMP" >> ~/.bash_profile
	fi
done

source ~/.bash_profile
