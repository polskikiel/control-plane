FROM golang:1.14.4-alpine3.12 as builder

ENV BASE_APP_DIR /go/src/github.com/kyma-project/control-plane/components/provisioner
WORKDIR ${BASE_APP_DIR}

ENV GO111MODULES=on

COPY ./go.mod ${BASE_APP_DIR}/go.mod
COPY ./go.sum ${BASE_APP_DIR}/go.sum

# Run go mod download first to take advantage of Docker caching
RUN apk add git && go mod download

COPY ./internal/ ${BASE_APP_DIR}/internal/
COPY ./pkg/ ${BASE_APP_DIR}/pkg/
COPY ./cmd/ ${BASE_APP_DIR}/cmd
COPY ./licenses ${BASE_APP_DIR}/licenses

RUN apk add -U --no-cache ca-certificates && update-ca-certificates

RUN go build -v -o main ./cmd/
RUN mkdir /app && mv ./main /app/main && mv ./licenses /app/licenses
RUN wget https://github.com/kyma-incubator/terraform-provider-gardener/releases/download/v0.0.3/terraform-provider-gardener-linux-amd64


FROM alpine:3.12.0
LABEL source = git@github.com:kyma-project/control-plane.git

# Hydroform needs write permissions
RUN mkdir /app && chmod 777 /app
WORKDIR /app

COPY --from=builder /app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/kyma-project/control-plane/components/provisioner/terraform-provider-gardener-linux-amd64 /.terraform.d/plugins/linux_amd64/terraform-provider-gardener

RUN chmod +x /.terraform.d/plugins/linux_amd64/terraform-provider-gardener

CMD ["/app/main"]
