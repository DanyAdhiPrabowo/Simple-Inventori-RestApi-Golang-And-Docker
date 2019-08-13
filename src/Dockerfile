FROM alpine
MAINTAINER Dany Adhi Prabowo <danyadhi4149@gmail.com>

#Working Directory
WORKDIR /home/app

#Copy in the source
COPY ./main ./main

#Make shell script executable
RUN chmod +x ./main

#Running
CMD ["./main"]