#!/bin/bash -eu

SRC=/python
DIST=/dist/layer.zip

pip3 install -t ${SRC} -U https://github.com/tetsuzawa/questplus/archive/master.zip
rm -f ${DIST}
ls -al .
zip -q -r ${DIST} ${SRC}