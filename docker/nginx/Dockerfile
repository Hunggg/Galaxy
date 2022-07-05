FROM nginx:alpine
RUN rm /etc/nginx/conf.d/*
COPY default.conf /etc/nginx/conf.d/default.conf
COPY nginx-cert.pem /etc/ssl/certs/nginx-cert.pem
COPY nginx-key.pem /etc/ssl/certs/nginx-key.pem
WORKDIR /home/
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 80
EXPOSE 443

ENTRYPOINT [ "./entrypoint.sh" ]
CMD ["/bin/sh"]