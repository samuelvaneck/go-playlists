CREATE TABLE IF NOT EXISTS radio_stations(
   id                serial PRIMARY KEY,
   name              varchar(255) NOT NULL,
   created_at        timestamp NOT NULL DEFAULT now(),
   updated_at        timestamp NOT NULL DEFAULT now(),
   api_url           varchar(255) NOT NULL,
   stream_url        varchar(255) NOT NULL,
   scraper_processor varchar(255) NOT NULL,
   scraper_import    boolean NOT NULL DEFAULT false
);
