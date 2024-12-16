#!/bin/bash

set -e

echo_info() {
    echo -e "\033[1;34m[INFO]\033[0m $1"
}

echo_warning() {
    echo -e "\033[1;33m[WARNING]\033[0m $1"
}

echo_error() {
    echo -e "\033[1;31m[ERROR]\033[0m $1" >&2
}

MIN_GO_VERSION="1.23"

version_greater_equal() {
   local IFS=.
    local i ver1=($1) ver2=($2)
    for ((i=${#ver1[@]}; i<${#ver2[@]}; i++)); do
        ver1[i]=0
    done
    for ((i=0; i<${#ver1[@]}; i++)); do
        if [[ -z ${ver2[i]} ]]; then
            ver2[i]=0
        fi
        if ((10#${ver1[i]} > 10#${ver2[i]})); then
            return 0
        elif ((10#${ver1[i]} < 10#${ver2[i]})); then
            return 1
        fi
    done
    return 0
}

compare_go_version() {
    INSTALLED_VERSION=$(go version 2>/dev/null | awk '{print $3}' | sed 's/go//')
    if version_greater_equal "$INSTALLED_VERSION" "$MIN_GO_VERSION"; then
        return 0
    else
        return 1
    fi
}

if command -v go >/dev/null 2>&1; then
    if compare_go_version; then
        echo_info "Go version ($(go version | awk '{print $3}')) is sufficient."
    else
        echo_error "Installed Go version ($(go version | awk '{print $3}')) is older than the required minimum version (go$MIN_GO_VERSION)."
        exit 1
    fi
else
    echo_error "Go is not installed on the system."
    exit 1
fi

BOB_DIR="$HOME/bob"
if [[ -d "$BOB_DIR" ]]; then
    echo_info "Directory $BOB_DIR already exists. Updating the repository..."
    cd "$BOB_DIR"
    git pull
else
    echo_info "Cloning the Bob repository..."
    git clone https://github.com/charmingruby/bob.git "$BOB_DIR"
fi

cd "$BOB_DIR"

echo_info "Building Bob..."
go build -o bob ./cmd/cli/main.go

echo_info "Installing Bob to /usr/local/bin..."
sudo mv bob /usr/local/bin/

CONFIG_DIR="$HOME/.config/bob"
mkdir -p "$CONFIG_DIR"

echo_info "Copying messages directory to $CONFIG_DIR..."
cp -r "$BOB_DIR/messages" "$CONFIG_DIR/"

echo_info "Installation complete! Verify by running 'bob'"

echo_info "Installation finished."