CREATE TABLE IF NOT EXISTS playlists(
   id                serial PRIMARY KEY,
   song_id           BIGINT NOT NULL,
   radio_station_id  BIGINT NOT NULL,
   created_at        timestamp NOT NULL DEFAULT now(),
   broadcasted_at    timestamp NOT NULL DEFAULT now(),
   scraper_import    boolean NOT NULL DEFAULT false
);
