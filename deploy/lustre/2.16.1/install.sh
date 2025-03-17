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
baseurl=https://downloads.whamcloud.com/public/lustre/lustre-2.16.1/el9.4/server
enabled=0
gpgcheck=0
EOF
    info "installing lustre from whamcloud public repo..."
    yum --disablerepo=* --enablerepo=lustre-server install lustre-zfs-dkms lustre-osd-zfs-mount lustre-iokit lustre lustre-resource-agents kmod-lustre -y > /dev/null 2>&1
    check_error "failed to install lustre packages"

info "lustre 2.16.1 installed successfully!"
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
    echo "| NOTE: This script automates the deployment of the Lustre 2.16.1 file system using ZFS as the storage backend. |"
    echo "================================================================================================================"
    echo ""

    ! grep -q "Rocky Linux release 9.4" /etc/redhat-release && error "os not match, this script is only for Rocky Linux release 9.4!
        download link: https://dl.rockylinux.org/vault/rocky/9.4/isos/x86_64/Rocky-9.4-x86_64-minimal.iso
    "

    [ "5.14.0-427.13.1.el9_4.x86_64" != $(uname -r) ] && error "kernel version(`uname -r`) not match, required kernel is 5.14.0-427.13.1.el9_4.x86_64!
        download link: https://dl.rockylinux.org/vault/rocky/9.4/isos/x86_64/Rocky-9.4-x86_64-minimal.iso
    "

    info "os info: `cat /etc/redhat-release` `uname -r` check passed..."

    online_install_lustre
}

main
