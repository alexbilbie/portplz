FROM scratch

COPY portplz /portplz
COPY test-metadata.json /test-metadata.json

ENV ECS_CONTAINER_METADATA_FILE=/test-metadata.json

ENTRYPOINT ["/portplz"]