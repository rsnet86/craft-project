FROM ubuntu
RUN apt update -y 
RUN apt install -y ca-certificates 
#RUN apk add --update tzdata
ENV TZ=Asia/Kolkata
## Clean Apk cache
#ERUN rm -rf /var/cache/apk/*
RUN mkdir /craft
WORKDIR /craft

COPY  url-shortener /craft

EXPOSE 8080
CMD ["./url-shortener"]
