
CREATE TABLE sellers(
    id bigserial NOT NULL,
    nama TEXT NOT NULL,
    hp varchar(30) NULL,
    alamat text null,
    status int default 0,
    last_update_by varchar(255) NULL,
	last_update bigint NULL,
    PRIMARY KEY (id)
);

create table customers ( 
    id bigserial NOT NULL,
    nama TEXT NOT NULL,
    hp varchar(30) NULL,
    alamat text null,
    coordinate text null,
    status int default 0,
    last_update_by varchar(255) NULL,
	last_update bigint NULL,
    PRIMARY KEY (id)    
);

create table drivers (  
    id bigserial NOT NULL,
    nama TEXT NOT NULL,
    hp varchar(30) NULL,
    alamat text null,
    photo text null,
    status int default 0,
    last_update_by varchar(255) NULL,
	last_update bigint NULL,
    PRIMARY KEY (id)    
);

create table admins ( 
    id bigserial NOT NULL,
    nama TEXT NOT NULL,
    status int default 0,
    last_update_by varchar(255) NULL,
	last_update bigint NULL,
    PRIMARY KEY (id)    
);

create table transaction_pickup ( 
    id bigserial not null ,
    
    jam_request_antar varchar(10),
    tanggal_request_antar bigint,
    nama_product TEXT NOT NULL,
    status int default 0,
    coordinate_tujuan text null,
   	keterangan text null,
   	
    jam_ambil varchar(10),
    tanggal_ambil bigint,
    photo_ambil text,
    
    photo_sampai text,
   	jam_sampai varchar(10),
   	tanggal_sampai bigint,
   	
   	id_seller bigserial ,
   	id_driver bigserial ,
   	id_customer bigserial ,
   	id_admin bigserial ,
   	
   	
    last_update_by varchar(255) NULL,
	last_update bigint NULL,
    PRIMARY KEY (id)    
);


