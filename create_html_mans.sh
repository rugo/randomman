#!/bin/bash

export IFS=":"
MANPAGEDIR="man_pages"

mkdir $MANPAGEDIR 

for p in $(manpath)
do
    cp -R $p/man? $MANPAGEDIR 2>/dev/null 
done

unset IFS

for file in $(find ${MANPAGEDIR}/man?/)
do
    zcat $file | groff -mandoc -Thtml > $file.html 
done
