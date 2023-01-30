#!/bin/bash

localarea=$1

ln -sf /usr/share/zoneinfo/$localarea /etc/localtime

hwclock --systohc
