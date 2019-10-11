create user ntci;
alter role ntci with password '123456';
create database ntci;
grant all privileges on database ntci to ntci;
alter database ntci owner to ntci;