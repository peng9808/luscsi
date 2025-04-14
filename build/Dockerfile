FROM rockylinux:8

RUN yum install nss -y

COPY ./_output/luscsi /

ENTRYPOINT [ "/luscsi" ]