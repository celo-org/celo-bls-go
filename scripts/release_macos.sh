#!/bin/bash -e

DIRECTORY=./libs
if [[ -d "$DIRECTORY" ]]
then
    echo "$DIRECTORY exists on your filesystem. Delete it and run the script again."
    exit 0
fi

pushd celo-bls-snark-rs/crates/bls-snark-sys

export RUSTFLAGS="-Ccodegen-units=1"
rustup default 1.52.1
cargo install cargo-strip -f

rustup target add x86_64-apple-darwin

cargo build --target=x86_64-apple-darwin --release -p bls-snark-sys
cargo strip --target x86_64-apple-darwin

popd

source `dirname $0`/copy_libs.sh