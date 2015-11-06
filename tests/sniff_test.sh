echo running tests for golang
UUID=$(cat /proc/sys/kernel/random/uuid)

pass "Unable to start the $VERSION container" docker run --privileged=true -d --name $UUID nanobox/build-golang sleep 365d

defer docker kill $UUID

pass "Unable to create code folder" docker exec $UUID mkdir -p /opt/code

fail "Detected something when there shouldn't be anything" docker exec $UUID bash -c "cd /opt/engines/golang/bin; ./sniff /opt/code"

pass "Failed to inject Gomfile" docker exec $UUID touch /opt/code/Gomfile

pass "Failed to detect" docker exec $UUID bash -c "cd /opt/engines/golang/bin; ./sniff /opt/code"