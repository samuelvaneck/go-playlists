CREATE TABLE IF NOT EXISTS artists_songs (
  artist_id BIGINT NOT NULL,
  song_id BIGINT NOT NULL,
  PRIMARY KEY (artist_id, song_id),
  FOREIGN KEY (artist_id) REFERENCES artists(id),
  FOREIGN KEY (song_id) REFERENCES songs(id)
);
