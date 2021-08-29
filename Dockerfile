FROM --platform=$BUILDPLATFORM golang:alpine AS build
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"
WORKDIR /root/
COPY main.go .
RUN go build main.go

FROM alpine
WORKDIR /root/
COPY --from=build /root/main .
EXPOSE 2701
CMD ["./main"]
