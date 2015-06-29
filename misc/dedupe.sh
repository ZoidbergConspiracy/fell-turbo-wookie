#!/bin/sh

HASHCMD="md5 -r"

LASTHASH=""
LASTFILE=""

for F in "$@"; do
  ${HASHCMD} "${F}" 2>&1
done | \
sort | \
while read LINE; do

  if [ "$( echo ${LINE} | grep -E '^[0-9a-f]{32}' )" ]; then
     HASH=$( echo ${LINE} | cut -d' ' -f 1 )
     FNAME=$( echo ${LINE} | cut -d' ' -f 2- )
     if [ "${HASH}" == "${LASTHASH}" ]; then
       echo "${FNAME} is a duplicate of ${LASTFILE}"
       rm -f "${FNAME}"
       ln -f "${LASTFILE}" "${FNAME}"
     else
       LASTHASH=${HASH}
       LASTFILE=${FNAME}
     fi
  fi

done

