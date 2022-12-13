#!/bin/sh

cd /root/git/repostats

git pull

# build Go app
go build -o bin/repostats .
echo 'Built Go app!'

# build Svelte Node app
cd /root/git/repostats/frontend
npm run build
echo 'Built Svelte Node app!'

# restart services
service repostats-api restart
service repostats-ssr restart
