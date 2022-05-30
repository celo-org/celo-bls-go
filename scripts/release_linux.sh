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
cargo install cargo-strip

rustup target add x86_64-unknown-linux-gnu

cargo build --target=x86_64-unknown-linux-gnu --release -p bls-snark-sys
cargo strip --target x86_64-unknown-linux-gnu

popd

source `dirname $0`/copy_libs.sh
