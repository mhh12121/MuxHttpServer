use lokcol;
CREATE TABLE IF NOT EXISTS int_auth_token_cache(
    tokenId int(64) auto_increment,
    id varchar(16) not null,
    int_auth_token varchar(16) unique key not null,
    device varchar(64),
    ip varchar(16),
    constraint primary key(tokenId)
)ENGINE= InnoDB DEFAULT charset=utf8;

INSERT INTO int_auth_token_cache(tokenId,id,int_auth_token,device,ip)
values(null,'sp100029','xxxyyyzzz','iPhone6','192.168.1.88'),
(null,'sp100029','aaabbbccc','Samsung Galaxy S3','177.15.3.8'),
(null,'sp100030','pppqqqsss','Samsung NOte pad','192.168.57.1'),
(null,'sp100031','dddeeefff','Xiao mi 5x','192.168.57.2'),
(null,'sp100031','eeefffggg','Xiao mi 4','111.20.3.7'),
(null,'sp100032','yuqbaJnmr','iPhoneSE','121.2.88.137')
