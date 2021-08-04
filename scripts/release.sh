#!/bin/bash -e

DIRECTORY=./libs
if [[ -d "$DIRECTORY" ]]
then
    echo "$DIRECTORY exists on your filesystem. Delete it and run the script again."
    exit 0
fi

if [[ ! -z $CIRCLE_SHA1 && $(git log --format=oneline -n 1 $CIRCLE_SHA1) != *"BUNDLE"* ]]; then
  echo "Not local and commit message doesn't contain BUNDLE, building only linux."
  source `dirname $0`/release_linux.sh
  exit 0
fi
pushd celo-bls-snark-rs/crates/bls-snark-sys

export RUSTFLAGS="-Ccodegen-units=1"
rustup default 1.52.1
cargo install cargo-lipo -f
cargo install cargo-strip -f

rustup target add i686-unknown-linux-gnu
rustup target add x86_64-unknown-linux-gnu
rustup target add arm-unknown-linux-gnueabi
rustup target add arm-unknown-linux-gnueabihf
rustup target add aarch64-unknown-linux-gnu
rustup target add mips-unknown-linux-gnu
rustup target add mipsel-unknown-linux-gnu
rustup target add mips64-unknown-linux-gnuabi64
rustup target add mips64el-unknown-linux-gnuabi64
rustup target add x86_64-apple-darwin
rustup target add aarch64-apple-darwin
rustup target add i686-pc-windows-gnu
rustup target add x86_64-pc-windows-gnu
rustup target add aarch64-linux-android
rustup target add armv7-linux-androideabi
rustup target add i686-linux-android
rustup target add x86_64-linux-android
rustup target add x86_64-unknown-linux-musl
rustup target add s390x-unknown-linux-gnu
rustup target add aarch64-apple-ios x86_64-apple-ios

cargo build --release --target=aarch64-linux-android --lib -p bls-snark-sys
cargo strip --target aarch64-linux-android
cargo build --release --target=armv7-linux-androideabi --lib -p bls-snark-sys
cargo strip --target armv7-linux-androideabi
cargo build --release --target=i686-linux-android --lib -p bls-snark-sys
cargo strip --target i686-linux-android
cargo build --release --target=x86_64-linux-android --lib -p bls-snark-sys
cargo strip --target x86_64-linux-android
cargo build --release -p bls-snark-sys
cargo strip
cargo build --target=i686-unknown-linux-gnu --release -p bls-snark-sys
cargo strip --target i686-unknown-linux-gnu
cargo build --target=x86_64-unknown-linux-gnu --release -p bls-snark-sys
cargo strip --target x86_64-unknown-linux-gnu
cargo build --target=arm-unknown-linux-gnueabi --release -p bls-snark-sys
cargo strip --target arm-unknown-linux-gnueabi
cargo build --target=arm-unknown-linux-gnueabihf --release -p bls-snark-sys
cargo strip --target arm-unknown-linux-gnueabihf
cargo build --target=aarch64-unknown-linux-gnu --release -p bls-snark-sys
cargo strip --target aarch64-unknown-linux-gnu
cargo build --target=mips-unknown-linux-gnu --release -p bls-snark-sys
cargo strip --target mips-unknown-linux-gnu
cargo build --target=mipsel-unknown-linux-gnu --release -p bls-snark-sys
cargo strip --target mipsel-unknown-linux-gnu
cargo build --target=mips64-unknown-linux-gnuabi64 --release -p bls-snark-sys
cargo strip --target mips64-unknown-linux-gnuabi64
cargo build --target=mips64el-unknown-linux-gnuabi64 --release -p bls-snark-sys
cargo strip --target mips64el-unknown-linux-gnuabi64
cargo build --target=x86_64-apple-darwin --release -p bls-snark-sys
cargo strip --target x86_64-apple-darwin
cargo build --target=aarch64-apple-darwin --release -p bls-snark-sys
cargo strip --target aarch64-apple-darwin
cargo build --target=i686-pc-windows-gnu --release -p bls-snark-sys
cargo strip --target i686-pc-windows-gnu
cargo build --target=x86_64-pc-windows-gnu --release -p bls-snark-sys
cargo strip --target x86_64-pc-windows-gnu
cargo build --target=x86_64-unknown-linux-musl --release -p bls-snark-sys
cargo strip --target x86_64-unknown-linux-musl
cargo build --target=s390x-unknown-linux-gnu --release -p bls-snark-sys
cargo strip --target s390x-unknown-linux-gnu
cargo lipo --release --targets=aarch64-apple-ios,x86_64-apple-ios -p bls-snark-sys

popd

source `dirname $0`/copy_libs.sh


tar czf libs.tar.gz libs
