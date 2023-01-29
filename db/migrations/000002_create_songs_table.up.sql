CREATE TABLE IF NOT EXISTS songs(
   id                   serial PRIMARY KEY,
   title                varchar(255) NOT NULL,
   artist               varchar(255) NOT NULL,
   isrc                 varchar(255) NOT NULL,
   id_on_spotify        varchar(255) NOT NULL,
   spotify_song_url     varchar(255) NOT NULL,
   spotify_artwork_url  varchar(255) NOT NULL,
   created_at           timestamp NOT NULL DEFAULT now(),
   updated_at           timestamp NOT NULL DEFAULT now(),
   full_name            varchar(255) NOT NULL
);
