cd ../  && git fetch --all &&  git reset --hard origin/master && godep restore &&  goconvey -port 8080 -host 192.168.1.134;
