FROM ubuntu:22.04

WORKDIR  /

ENV DEBIAN_FRONTEND=noninteractive \
    GOROOT=/opt/go \
    GOPATH=/root/.go

# Install packagesr equired for vscode, gstreamer runtime and goc gompiling
RUN set -x; apt-get update && \
    apt-get -y install gdb wget pkg-config glib-2.0 gcc libsoup2.4-1 libpango1.0 git curl vim net-tools bash-completion git-lfs build-essential swig unzip && \
    echo "source /etc/profile.d/bash_completion.sh" >> ~/.bashrc && \
    apt-get -y autoremove

# Install GO
ARG GO_VER="1.21.11"

RUN ARCH=$(dpkg --print-architecture) && \
    set -x; cd / && \
    mkdir $GOPATH && \
    cd /opt && wget -q https://storage.googleapis.com/golang/go${GO_VER}.linux-${ARCH}.tar.gz && \
    tar zxf go${GO_VER}.linux-${ARCH}.tar.gz && rm go${GO_VER}.linux-${ARCH}.tar.gz && \
    ln -s /opt/go/bin/go /usr/bin/ && \
    apt-get -y autoremove

# Install some go tools for vscode to avoid it prompting to install them
RUN set -x; go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    export GO111MODULE=on && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest  && \
    go install golang.org/x/tools/gopls@latest



