DROP TABLE IF EXISTS articles;

CREATE SEQUENCE article_id_seq;

CREATE TABLE articles (
    id integer NOT NULL DEFAULT nextval('article_id_seq'),
	article_id text,
    title text,
	description text,
	link text,
	published date,
	author text
);

ALTER TABLE ONLY articles
    ADD CONSTRAINT articles_pkey PRIMARY KEY (id);