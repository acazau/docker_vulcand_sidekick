FROM scratch

COPY docker_vulcand_sidekick /
ENTRYPOINT ["/docker_vulcand_sidekick"]