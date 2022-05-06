-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE sellers(
    id int NOT NULL AUTO_INCREMENT,
    nama TEXT NOT NULL,
    hp varchar(30) NULL,
    kode varchar(6) NULL,
    password varchar(25) NULL,
    alamat text null,
    status int default 0,
    last_update_by varchar(255) NULL,
	last_update bigint NULL,
    PRIMARY KEY (id)
);
CREATE INDEX idx_kode ON sellers  (kode);


create table customers ( 
    id int NOT NULL AUTO_INCREMENT,
    seller_id bigint ,
    nama TEXT NOT NULL,
    hp varchar(30) NULL,
    alamat text null,
    coordinate text null,
    status int default 0,
    last_update_by varchar(255) NULL,
	last_update bigint NULL,
    PRIMARY KEY (id)    
);
CREATE INDEX idx_seller ON customers (seller_id);

create table drivers (  
    id int NOT NULL AUTO_INCREMENT,
    nama varchar(100) NOT NULL,
    kode varchar(6) NULL,
    password varchar(25) NULL,
    hp varchar(30) NULL,
    alamat text null,
    photo text null,
    status int default 0,
    last_update_by varchar(255) NULL,
	last_update bigint NULL,
    PRIMARY KEY (id)    
);
CREATE INDEX idx_kode ON drivers (kode);
CREATE INDEX idx_nama ON drivers (nama);

create table admins ( 
    id int NOT NULL AUTO_INCREMENT,
    kode varchar(6) NULL,
    nama TEXT NOT NULL,
    password varchar(20) NULL,
    status int default 0,
    last_update_by varchar(255) NULL,
	last_update bigint NULL,
    role_id int(11) DEFAULT NULL,
    PRIMARY KEY (id)    
);
CREATE INDEX idx_kode ON admins  (kode);


CREATE TABLE uruts(
    id int NOT NULL AUTO_INCREMENT,
    prefix varchar(10)  NOT NULL,
    keterangan varchar(30) NULL,
    no_terakhir bigint  default 0,
    PRIMARY KEY (id)
);
CREATE INDEX idx_prefix ON uruts  (prefix);


create table transaksi ( 
    id int NOT NULL AUTO_INCREMENT,
    transaksi_date int null,
    tanggal_request_antar int,
    jam_request_antar varchar(10),
    nama_product varchar (200) NOT NULL,
    status int default 0,
    coordinate_tujuan text null,
   	keterangan text null,
    photo_ambil text,
    tanggal_ambil_str varchar(20),
    tanggal_ambil int,
    photo_sampai text,
   	tanggal_sampai_str varchar(20),
   	tanggal_sampai int,
   	
   	id_seller int ,
   	id_driver int ,
   	id_customer int ,
   	id_admin int ,
   	
    last_update_by varchar(255) NULL,
	last_update int NULL,
    PRIMARY KEY (id)    
);
CREATE INDEX idx_trx_date ON transaksi  (transaksi_date);
CREATE INDEX idx_tanggal_request_antar ON transaksi  (tanggal_request_antar);
CREATE INDEX idx_seller ON transaksi  (id_seller);
CREATE INDEX idx_driver ON transaksi  (id_driver);


CREATE TABLE m_menus (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    description character varying(100) NOT NULL,
    last_update_by character varying(100),
    last_update timestamp without time zone,
    link character varying(200),
    parent_id bigint,
    icon character varying(50),
    status integer
);


CREATE TABLE m_role_menu (
    role_id bigint NOT NULL,
    menu_id bigint NOT NULL,
    status integer,
    last_update_by character varying(100),
    last_update timestamp without time zone
);

CREATE TABLE public.m_roles (
    id bigint NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(255) NOT NULL,
    last_update_by character varying(100),
    last_update timestamp without time zone
);


CREATE TABLE public.m_role_user (
    role_id bigint NOT NULL,
    user_id bigint NOT NULL,
    last_update_by character varying(100),
    last_update timestamp without time zone,
    status integer DEFAULT 0
);



-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

