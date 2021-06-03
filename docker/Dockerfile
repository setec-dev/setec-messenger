FROM alpine:latest

# deps
RUN apk add --no-cache git go


# configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
  
# install steps
RUN 
RUN git clone https://github.com/setec-dev/setec-messenger.git

