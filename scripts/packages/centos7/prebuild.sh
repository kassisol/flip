#!/usr/bin/env bash

ROOTDIR=$(dirname $0)/../../..
cd $(dirname $0)

if [ -d "build" ]; then
	rm -rf build
fi
mkdir -p build

cp ${ROOTDIR}/contrib/init/systemd/flip.service build/
cp ${ROOTDIR}/bin/flip build/
