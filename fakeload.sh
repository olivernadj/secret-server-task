#!/usr/bin/env bash

################################################################
# Script name      : fakeload.sh                               #
# Description      : Makes randomised apache bench requests    #
# Original Author  : Oliver Nadj <mr.oliver.nadj@gmail.com>    #
################################################################


# validate the arguments
if [[ -z "$1" ]]; then HOST="http://localhost:8080/v1"; else HOST=$1; fi


loadGet () {
    if [[ -z "$2" ]]; then SLEEP_RAND=5; else SLEEP_RAND=$2; fi
    if [[ -z "$3" ]]; then REQUESTS_RAND=5; else REQUESTS_RAND=$3; fi
    (
        while [[ 1 ]]   # Endless loop.
        do
           sleep $(( $RANDOM % $SLEEP_RAND + 1 ))
           REQUESTS=$(( $RANDOM % REQUESTS_RAND + 1 ))
           echo "ab -n ${REQUESTS} $1"
           `ab -n ${REQUESTS} $1 > /dev/null`
        done
    ) &
}

loadPost () {
    if [[ -z "$2" ]]; then SLEEP_RAND=5; else SLEEP_RAND=$2; fi
    if [[ -z "$3" ]]; then REQUESTS_RAND=5; else REQUESTS_RAND=$3; fi
    (
        while [[ 1 ]]   # Endless loop.
        do
           sleep $(( $RANDOM % $SLEEP_RAND + 1 ))
           REQUESTS=$(( $RANDOM % REQUESTS_RAND + 1 ))
           echo "ab -p fakeload-post.txt -T application/x-www-form-urlencoded -n ${REQUESTS} $1"
           `ab -p fakeload-post.txt -T application/x-www-form-urlencoded -n ${REQUESTS} $1> /dev/null`
        done
    ) &
}

loadGet "$HOST/" 7 50
loadPost "$HOST/secret" 20 100
loadGet "$HOST/secret/test" 15 300


trap ctrl_c INT

function ctrl_c() {
        echo "Bye"
        pkill -P $$
        exit 1
}

while [[ 1 ]]
do
    sleep 1
done

