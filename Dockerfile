FROM golang:1.15
WORKDIR /go/src/os_snmp
COPY app.yaml .
COPY os_snmp .
COPY conf ./conf
COPY script ./script
EXPOSE 8800
#RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN go build -o os_snmp
CMD ["/bin/bash", "/go/src/os_snmp/script/build.sh"]