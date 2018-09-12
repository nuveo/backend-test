CREATE USER nuveo WITH
	LOGIN
	NOSUPERUSER
	NOCREATEDB
	NOCREATEROLE
	INHERIT
	NOREPLICATION
	CONNECTION LIMIT -1
	PASSWORD 'nuveo';

CREATE DATABASE nuveo
    WITH
    OWNER = nuveo
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;