## wallet core #######################################################
FROM ubuntu:18.04 AS walletcore

RUN apt-get update \
    && apt-get install -y \
        wget \
        curl \
        git \
        vim \
        unzip \
        xz-utils \
        software-properties-common \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN wget -O - https://apt.kitware.com/keys/kitware-archive-latest.asc | apt-key add - \
    && apt-add-repository 'deb https://apt.kitware.com/ubuntu/ bionic main' \
    && apt-add-repository -y ppa:mhier/libboost-latest

RUN apt-get update \
    && apt-get install -y \
        build-essential \
        libtool autoconf pkg-config \
        ninja-build \
        ruby-full \
        clang-10 \
        llvm-10 \
        libc++-dev libc++abi-dev \
        cmake \
        libboost1.74-dev \
        ccache \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

ENV CC=/usr/bin/clang-10
ENV CXX=/usr/bin/clang++-10

RUN git clone https://github.com/ttmbank/wallet-core.git
WORKDIR /wallet-core
RUN git checkout master-x

RUN tools/install-dependencies

RUN tools/generate-files \
    && cmake -H. -Bbuild -DCMAKE_BUILD_TYPE=Debug \
    && make -Cbuild -j12


## server builder #######################################################
FROM golang:1.14 AS build

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /tmp/app

COPY . .
COPY --from=walletcore /wallet-core/include /wallet-core/include
COPY --from=walletcore /wallet-core/build /wallet-core/build


RUN GOOS=linux go build -o ./backend/server ./app/app.go

## server #######################################################
FROM ubuntu:18.04 AS server

RUN apk add ca-certificates

COPY --from=build /tmp/app/backend/server /app/server
COPY --from=walletcore /wallet-core/include /app/wallet-core/include
COPY --from=walletcore /wallet-core/build/libprotobuf.a /app/wallet-core/build/libprotobuf.a
COPY --from=walletcore /wallet-core/build/libTrustWalletCore.a /app/wallet-core/build/libTrustWalletCore.a

WORKDIR "/app"

EXPOSE 8020

CMD ["./server"]

