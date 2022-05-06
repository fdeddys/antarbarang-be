-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied


CREATE TABLE public.parameter (
	id bigserial NOT NULL,
	"name" varchar(255) NULL,
	value varchar(255) NULL,
	isviewable int2 NULL DEFAULT 1,
	last_update_by varchar(255) NULL,
	last_update bigint NULL,
	PRIMARY KEY (id)
);

INSERT INTO public."parameter" ("name", value, isviewable) values ('tax', '10', 1);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back