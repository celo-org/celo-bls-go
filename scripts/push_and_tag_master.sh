#!/bin/bash -ex

TAG=$1

source `dirname $0`/push_and_tag_lib.sh
push_tag $TAG
