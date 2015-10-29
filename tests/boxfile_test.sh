echo running tests for golang
UUID=$(cat /proc/sys/kernel/random/uuid)

pass "Unable to start the $VERSION container" docker run --privileged=true -d --name $UUID nanobox/build-golang sleep 365d

defer docker kill $UUID

pass "Failed to execute the boxfile script" docker exec $UUID bash -c "cd /opt/engines/golang/bin; ./boxfile '$(payload default-boxfile)'"
