FROM golang:1.23.2-alpine AS base
WORKDIR /app

# ------------------------------------------------

FROM base AS builder

COPY . .

RUN apk add --no-cache gcc musl-dev

RUN wget -qO- https://github.com/go-task/task/releases/download/v3.39.2/task_linux_amd64.tar.gz | tar xvz -C /usr/local/bin

RUN task build-plugins
RUN task build-app
RUN task create-configs

# ------------------------------------------------

FROM base AS runner

RUN apk add --no-cache envsubst

RUN mkdir /configs
RUN mkdir /plugins
RUN mkdir /certs

COPY --from=builder /app/.configs/* /configs
COPY --from=builder /app/.certs/* /certs
COPY --from=builder /plugins/* /plugins
COPY --from=builder /app/entrypoint/seam ./seam

RUN <<EOR
  set -e
  for dir in /configs/*; do
    cp $dir $dir.temp;
    cat $dir | envsubst > $dir.temp;
    mv $dir.temp $dir;
  done
EOR

RUN adduser -D seam
RUN chown -R seam:seam ./
RUN chmod +x ./seam

USER seam

CMD ["./seam"]
