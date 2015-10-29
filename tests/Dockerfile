FROM nanobox/build

# Remove existing engines, we don't want to test those
RUN rm -rf /opt/engines/*

# Create folders for this engine
RUN mkdir -p /opt/engines/golang/bin && \
    mkdir -p /opt/engines/golang/lib && \
    mkdir -p /opt/engines/golang/golangs

# Add in this engine's files
ADD bin/. /opt/engines/golang/bin
ADD lib/. /opt/engines/golang/lib
ADD golangs/. /opt/engines/golang/golangs