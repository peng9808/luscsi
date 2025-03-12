#!/bin/bash
# install-os-deps.sh is a helper to utility to install required packages for lustre
#

#set -e

DEBUG=1
function error
{
  echo -e "ERROR: \t$1"
  exit 1
}

function info
{
  echo -e "INFO: \t$1"
}

function debug
{
  if [ $DEBUG -ne 0 ]
  then
    echo -e "DEBUG: \t$1"
  fi
}

function check_error
{    
    if [ ! $? -eq 0 ]
    then
        error "$@"
    fi
}

function online_install_lustre
{
    tee /etc/yum.repos.d/lustre-server.repo <<EOF > /dev/null 2>&1
[lustre-server]
name=lustre-server
baseurl=https://downloads.whamcloud.com/public/lustre/lustre-2.15.4/el8.9/server
enabled=0
gpgcheck=0
EOF
    info "installing lustre from whamcloud public repo..."
    yum --disablerepo=* --enablerepo=lustre-server install lustre-zfs-dkms lustre-osd-zfs-mount lustre-iokit lustre lustre-resource-agents -y > /dev/null 2>&1
    check_error "failed to install lustre packages"

info "lustre 2.15.4 installed successfully!"
    show_lustre_version
}

function show_lustre_version
{
     lctl lustre_build_version

     modprobe zfs && zfs --version
}

function main
{
    echo "================================================================================================================"
    echo "| NOTE: This script automates the deployment of the Lustre 2.15.4 file system using ZFS as the storage backend. |"
    echo "================================================================================================================"
    echo ""

    ! grep -q "Rocky Linux release 8.9" /etc/redhat-release && error "os not match, this script is only for Rocky Linux release 8.9!
        download link: https://dl.rockylinux.org/vault/rocky/8.9/isos/x86_64/Rocky-8.9-x86_64-minimal.iso
    "

    [ "4.18.0-513.5.1.el8_9.x86_64" != $(uname -r) ] && error "kernel version(`uname -r`) not match, required kernel is 4.18.0-513.5.1.el8_9.x86_64!
        download link: https://dl.rockylinux.org/vault/rocky/8.9/isos/x86_64/Rocky-8.9-x86_64-minimal.iso
    "

    info "os info: `cat /etc/redhat-release` `uname -r` check passed..."

    online_install_lustre
}

main
