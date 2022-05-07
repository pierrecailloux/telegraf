FROM debian 
USER root 
RUN apt update 
RUN apt upgrade -y 
RUN apt install wget git -y 
RUN wget https://go.dev/dl/go1.18.1.linux-amd64.tar.gz
RUN  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz
RUN wget https://dl.influxdata.com/telegraf/releases/telegraf_1.22.3-1_amd64.deb
RUN  dpkg -i telegraf_1.22.3-1_amd64.deb
RUN rm  telegraf_1.22.3-1_amd64.deb
ENV PATH /usr/local/go/bin:$PATH
RUN mkdir -p /tmp/telegraf/ 
ADD   . /tmp/telegraf
RUN mkdir -p /etc/telegraf/plugins/Conso
WORKDIR /tmp/telegraf
RUN go build  -o /etc/telegraf/plugins/Conso/conso  /tmp/telegraf/cmd/main.go
COPY  telegraf.conf /etc/telegraf/telegraf.conf
COPY  plugin.conf  /etc/telegraf/plugins/Conso/
CMD [ "telegraf" ]

