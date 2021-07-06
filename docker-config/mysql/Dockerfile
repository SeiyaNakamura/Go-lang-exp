FROM mysql

EXPOSE 3306

RUN apt-get update
RUN apt-get -y install locales-all

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8

COPY ./docker-config/mysql/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]