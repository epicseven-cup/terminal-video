FROM gocv/opencv:4.10.0-ubuntu-22.04
COPY . .
RUN wget https://go.dev/dl/go1.23.3.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz
RUN apt update && apt install jp2a -y
ENV PATH=$PATH:/usr/local/go/bin
RUN go build .
CMD ["sleep", "infinity"]