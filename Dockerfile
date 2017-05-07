FROM scratch
MAINTAINER mangirdas@judeikis.lt
COPY       api /bin/api
ENTRYPOINT ["/bin/api"]