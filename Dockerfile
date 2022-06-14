FROM --platform=${TARGETPLATFORM} alpine:3 AS executor

# copy over the binary from the first stage
COPY godent /godent/godent

WORKDIR "/godent"
ENTRYPOINT [ "/godent/godent" ]