FROM debian:stretch-slim
RUN useradd -u 10001 runner

FROM scratch

ENV SERVICE_NAME="CorrId"

WORKDIR /app
ADD bin/corr-id-generator /app/
ADD settings/local.yml /app/

COPY --from=0 /etc/passwd /etc/passwd
USER runner

ENTRYPOINT ["/app/corr-id-generator", "-config-type=local", "-settings=local.yml"]