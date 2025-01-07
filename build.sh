#!/bin/bash

cd $(dirname "$0") 
workdir=$(pwd)
echo $workdir

rm -rf "$workdir/dist"
mkdir "$workdir/dist"

go build -o "$workdir/dist/slidegoose" "$workdir/cmd/main.go" 

cd "$workdir/web"
pnpm build --outDir "$workdir/dist/web"
cd "$workdir"

./dist/slidegoose