-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied


INSERT INTO public.uruts (prefix, keterangan, no_terakhir) VALUES('S', 'Seller', 0);

INSERT INTO public.uruts (prefix, keterangan, no_terakhir) VALUES('D', 'Driver', 0);

INSERT INTO public.uruts (prefix, keterangan, no_terakhir) VALUES('A', 'Admin', 0);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back