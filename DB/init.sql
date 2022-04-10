DROP TABLE IF EXISTS article;

CREATE SEQUENCE article_id_seq;

CREATE TABLE article (
    id integer NOT NULL DEFAULT nextval('article_id_seq'),
	article_id text,
    title text,
	description text,
	link text,
	published date,
	author text
);

ALTER TABLE ONLY article
    ADD CONSTRAINT article_pkey PRIMARY KEY (id);

DROP TABLE IF EXISTS source;

CREATE SEQUENCE source_id_seq;

CREATE TABLE source (
    id integer NOT NULL DEFAULT nextval('source_id_seq'),
	url text
);

ALTER TABLE ONLY source
    ADD CONSTRAINT source_pkey PRIMARY KEY (id);

COPY source (url) FROM stdin;
http://feeds.bbci.co.uk/news/uk/rss.xml
http://feeds.bbci.co.uk/news/technology/rss.xml
http://feeds.skynews.com/feeds/rss/uk.xml
http://feeds.skynews.com/feeds/rss/technology.xml
\.
