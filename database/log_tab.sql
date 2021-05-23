CREATE TABLE public.info_log
(
    id SERIAL PRIMARY KEY,
   	info VARCHAR(250)
);

INSERT INTO public.info_log(info)
	VALUES 
	('log 1'),
    ('log 2'),
    ('log 3');