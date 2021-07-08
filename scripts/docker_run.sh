# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# OEML Configuration Manual
# https://docs.coinapi.io/oeml.html#configuration-parameters


# docker run -d --name=BINANCE --restart=always --net=host coinapi/oeml-api --env-file oeml.env

# docker run -d --restart=always --net=host coinapi/oeml-webui

# docker run -d --name=BINANCE --restart=always --net=host coinapi/oeml-api --env-file oeml.env --OEML:ExchangeId BINANCE

# one instance per exchange
# docker run -d --name=HITBTC --restart=always --net=host coinapi/oeml-api --env-file oeml.env --OEML:ExchangeId HITBTC

