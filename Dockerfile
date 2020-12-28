FROM golang:1.15
WORKDIR /app/go/os_snmp
COPY os_snmp .
COPY conf ./conf
COPY script ./script
EXPOSE 8800
#RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN go build -o os_snmp
CMD ["/bin/bash", "/app/go/os_snmp/script/build.sh"]