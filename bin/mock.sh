#!/bin/bash
set -ex

mockgen -source client.go  -destination mock/client.go -package mock
