#!/bin/bash -e

DIRECTORY=./libs

TOOLS_DIR=`dirname $0`
COMPILE_DIR=${TOOLS_DIR}/../celo-bls-snark-rs/target
rm -rf $COMPILE_DIR/x86_64-apple-ios $COMPILE_DIR/aarch64-apple-ios
for platform in `ls ${COMPILE_DIR} | grep -v release | grep -v debug`
do
  PLATFORM_DIR=${DIRECTORY}/$platform
  mkdir -p ${PLATFORM_DIR}
  LIB_PATH=${COMPILE_DIR}/$platform/release/libbls_snark_sys.a
  if [[ -f ${LIB_PATH} ]]
  then
    cp ${COMPILE_DIR}/$platform/release/libbls_snark_sys.a ${PLATFORM_DIR}
  fi
  WINDOWS_LIB_PATH=${COMPILE_DIR}/$platform/release/bls_snark_sys.lib
  if [[ -f ${WINDOWS_LIB_PATH} ]]
  then
    cp ${COMPILE_DIR}/$platform/release/bls_snark_sys.lib ${PLATFORM_DIR}
  fi
done