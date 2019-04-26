use lokcol;
CREATE TABLE IF NOT EXISTS simpleuser(
    id int(64) auto_increment,
    firstname varchar(16),
    lastname varchar(16),
    age int(3),
    constraint primary key (id)
)ENGINE=InnoDB DEFAULT charset=utf8;



INSERT INTO simpleuser(id,firstname,lastname,age) VALUES(null,"first1","last1",12),(null,"first2","last2",22),(null,"first3","last3",32)
