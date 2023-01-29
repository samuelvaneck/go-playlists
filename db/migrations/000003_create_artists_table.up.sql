CREATE TABLE IF NOT EXISTS artists(
   id                   serial PRIMARY KEY,
   name                 varchar(255) NOT NULL,
   id_on_spotify        varchar(255) NOT NULL,
   spotify_artist_url   varchar(255) NOT NULL,
   spotify_artwork_url  varchar(255) NOT NULL,
   created_at           timestamp NOT NULL DEFAULT now(),
   updated_at           timestamp NOT NULL DEFAULT now()
);
