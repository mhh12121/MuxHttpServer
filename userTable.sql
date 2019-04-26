use lokcol;
create table if not exists user(
    sid int(64) auto_increment,
    id varchar(16) unique not null,
    name varchar(32) not null,
    age int(3) not null,
    constraint primary key(sid)

)ENGINE=InnoDB DEFAULT charset=utf8;

INSERT INTO user(sid,id,name,age) VALUES(null,'sp100029','peter',12),(null,'sp100030','tom',25),(null,'sp100031','karl',32),(null,'sp100032','mary',40);
