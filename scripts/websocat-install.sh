# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# https://github.com/vi/websocat
# https://gist.github.com/mingliangguo/635345dcd6b603da337d4c71792bd330
# libssl-dev and pkg-config are required in order to build websocat
command sudo apt-get update

command sudo apt-get install -y curl build-essential libssl-dev pkg-config
command curl https://sh.rustup.rs -sSf | sh

export PATH=$HOME/.cargo/bin:$PATH

command cargo install --features=ssl websocat