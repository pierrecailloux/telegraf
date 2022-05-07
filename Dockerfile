FROM debian 
USER root 
RUN apt update 
RUN apt upgrade -y 
RUN apt install wget -y 
RUN wget https://go.dev/dl/go1.18.1.linux-amd64.tar.gz
RUN  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz
RUN wget https://dl.influxdata.com/telegraf/releases/telegraf_1.22.3-1_amd64.deb
RUN  dpkg -i telegraf_1.22.3-1_amd64.deb
RUN rm  telegraf_1.22.3-1_amd64.deb
ENV PATH /usr/local/go/bin:$PATH
RUN mkdir -p /tmp/telegraf/cmd 
RUN mkdir -p /tmp/telegraf/conso 
COPY cmd /tmp/telegraf/cmd 
COPY plugins/input/conso /tmp/telegraf/conso
RUN go build  /tmp/telegraf/co
RUN mkdir -p /etc/telegraf/plugins/Conso
COPY  
RUN go version 

