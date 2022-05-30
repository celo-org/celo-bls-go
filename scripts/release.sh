#!/bin/bash -e

DIRECTORY=./libs
if [[ -d "$DIRECTORY" ]]
then
    echo "$DIRECTORY exists on your filesystem. Delete it and run the script again."
    exit 1
fi

pushd celo-bls-snark-rs/crates/bls-snark-sys

export RUSTFLAGS="-Ccodegen-units=1"
rustup default 1.52.1
cargo install cargo-lipo -f
cargo install cargo-strip -f

declare -a platforms=("i686-unknown-linux-gnu" "x86_64-unknown-linux-gnu" "arm-unknown-linux-gnueabi" "arm-unknown-linux-gnueabihf" "aarch64-unknown-linux-gnu" "mips-unknown-linux-gnu" "mipsel-unknown-linux-gnu" "mips64-unknown-linux-gnuabi64" "mips64el-unknown-linux-gnuabi64" "x86_64-apple-darwin" "aarch64-apple-darwin" "i686-pc-windows-gnu" "x86_64-pc-windows-gnu" "aarch64-linux-android" "armv7-linux-androideabi" "i686-linux-android" "x86_64-linux-android" "x86_64-unknown-linux-musl" "s390x-unknown-linux-gnu" "aarch64-apple-ios x86_64-apple-ios" )

for platform in ${platforms[@]}; do
  rustup target add $platform
  cargo build --release --target=$platform --lib -p bls-snark-sys
  cargo strip --target $platform
done
cargo lipo --release --targets=aarch64-apple-ios,x86_64-apple-ios -p bls-snark-sys

popd

source `dirname $0`/copy_libs.sh


tar czf libs.tar.gz libs
