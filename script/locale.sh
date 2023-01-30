#!/bin/bash

locale=$1

echo $locale > /etc/locale.gen

locale-gen

echo "LANG=$locale" > /etc/locale.conf
