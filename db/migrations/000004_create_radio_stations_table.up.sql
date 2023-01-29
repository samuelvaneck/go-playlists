CREATE TABLE IF NOT EXISTS radio_stations(
   id                serial PRIMARY KEY,
   name              varchar(255) NOT NULL,
   created_at        timestamp NOT NULL DEFAULT now(),
   updated_at        timestamp NOT NULL DEFAULT now(),
   api_url           text NOT NULL,
   stream_url        varchar(255) NOT NULL,
   scraper_processor varchar(255) NOT NULL,
   scraper_import    boolean NOT NULL DEFAULT false
);

INSERT INTO radio_stations
  (name, api_url, stream_url, scraper_processor, scraper_import)
VALUES
  ('Radio 2', 'https://www.nporadio2.nl/api/tracks', 'https://icecast.samuelvaneck.com/radio2.mp3', 'npo_api_processor', true),
  ('Qmusic', 'https://api.qmusic.nl/2.4/tracks/plays?limit=1&next=true', 'https://icecast.samuelvaneck.com/qmusic.mp3', 'qmusic_api_processor', true),
  ('Sublime FM', 'https://sublime.nl/sublime-playlist/', 'https://icecast.samuelvaneck.com/sublimefm.mp3', 'scraper', true),
  ('Radio 5','https://www.nporadio5.nl/api/tracks', 'https://icecast.samuelvaneck.com/radio5.mp3', 'npo_api_processor', true),
  ('Radio 3FM', 'https://www.npo3fm.nl/api/tracks', 'https://icecast.samuelvaneck.com/radio3fm.mp3', 'npo_api_processor', true),
  ('Groot Nieuws Radio', 'https://www.grootnieuwsradio.nl/muziek/playlist', 'https://icecast.samuelvaneck.com/gnr.mp3', 'scraper', true),
  ('Radio 538', 'https://graph.talparad.io/?query=%7B%0A%20%20getStation(profile%3A%20%22radio-brand-web%22%2C%20slug%3A%20%22radio-538%22)%20%7B%0A%20%20%20%20title%0A%20%20%20%20playouts(profile%3A%20%22%22%2C%20limit%3A%2010)%20%7B%0A%20%20%20%20%20%20broadcastDate%0A%20%20%20%20%20%20track%20%7B%0A%20%20%20%20%20%20%20%20id%0A%20%20%20%20%20%20%20%20title%0A%20%20%20%20%20%20%20%20artistName%0A%20%20%20%20%20%20%20%20isrc%0A%20%20%20%20%20%20%20%20images%20%7B%0A%20%20%20%20%20%20%20%20%20%20type%0A%20%20%20%20%20%20%20%20%20%20uri%0A%20%20%20%20%20%20%20%20%20%20__typename%0A%20%20%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20%20%20__typename%0A%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20__typename%0A%20%20%20%20%7D%0A%20%20%20%20__typename%0A%20%20%7D%0A%7D%0A&variables=%7B%7D', 'https://icecast.samuelvaneck.com/radio538.mp3', 'talpa_api_processor', true),
  ('Sky Radio', 'https://graph.talparad.io/?query=%7B%0A%20%20getStation(profile%3A%20%22radio-brand-web%22%2C%20slug%3A%20%22sky-radio%22)%20%7B%0A%20%20%20%20title%0A%20%20%20%20playouts(profile%3A%20%22%22%2C%20limit%3A%2010)%20%7B%0A%20%20%20%20%20%20broadcastDate%0A%20%20%20%20%20%20track%20%7B%0A%20%20%20%20%20%20%20%20id%0A%20%20%20%20%20%20%20%20title%0A%20%20%20%20%20%20%20%20artistName%0A%20%20%20%20%20%20%20%20isrc%0A%20%20%20%20%20%20%20%20images%20%7B%0A%20%20%20%20%20%20%20%20%20%20type%0A%20%20%20%20%20%20%20%20%20%20uri%0A%20%20%20%20%20%20%20%20%20%20__typename%0A%20%20%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20%20%20__typename%0A%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20__typename%0A%20%20%20%20%7D%0A%20%20%20%20__typename%0A%20%20%7D%0A%7D%0A&variables=%7B%7D', 'https://icecast.samuelvaneck.com/skyradio.mp3', 'talpa_api_processor', true),
  ('Radio 1', 'https://www.nporadio1.nl/api/tracks', 'https://icecast.samuelvaneck.com/radio1.mp3', 'npo_api_processor', true),
  ('Radio 10', 'https://graph.talparad.io/?query=%7B%0A%20%20getStation(profile%3A%20%22radio-brand-web%22%2C%20slug%3A%20%22radio-10%22)%20%7B%0A%20%20%20%20title%0A%20%20%20%20playouts(profile%3A%20%22%22%2C%20limit%3A%2010)%20%7B%0A%20%20%20%20%20%20broadcastDate%0A%20%20%20%20%20%20track%20%7B%0A%20%20%20%20%20%20%20%20id%0A%20%20%20%20%20%20%20%20title%0A%20%20%20%20%20%20%20%20artistName%0A%20%20%20%20%20%20%20%20isrc%0A%20%20%20%20%20%20%20%20images%20%7B%0A%20%20%20%20%20%20%20%20%20%20type%0A%20%20%20%20%20%20%20%20%20%20uri%0A%20%20%20%20%20%20%20%20%20%20__typename%0A%20%20%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20%20%20__typename%0A%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20__typename%0A%20%20%20%20%7D%0A%20%20%20%20__typename%0A%20%20%7D%0A%7D%0A&variables=%7B%7D', 'https://icecast.samuelvaneck.com/radio10.mp3', 'talpa_api_processor', true),
  ('Radio Veronica', 'https://graph.talparad.io/?query=%7B%0A%20%20getStation(profile%3A%20%22radio-brand-web%22%2C%20slug%3A%20%22radio-veronica%22)%20%7B%0A%20%20%20%20title%0A%20%20%20%20playouts(profile%3A%20%22%22%2C%20limit%3A%2010)%20%7B%0A%20%20%20%20%20%20broadcastDate%0A%20%20%20%20%20%20track%20%7B%0A%20%20%20%20%20%20%20%20id%0A%20%20%20%20%20%20%20%20title%0A%20%20%20%20%20%20%20%20artistName%0A%20%20%20%20%20%20%20%20isrc%0A%20%20%20%20%20%20%20%20images%20%7B%0A%20%20%20%20%20%20%20%20%20%20type%0A%20%20%20%20%20%20%20%20%20%20uri%0A%20%20%20%20%20%20%20%20%20%20__typename%0A%20%20%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20%20%20__typename%0A%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20__typename%0A%20%20%20%20%7D%0A%20%20%20%20__typename%0A%20%20%7D%0A%7D%0A&variables=%7B%7D', 'https://icecast.samuelvaneck.com/veronica.mp3', 'talpa_api_processor', true),
  ('SLAM!', 'https://api.slam.nl/api/live?brand=slam', 'https://icecast.samuelvaneck.com/slam.mp3', 'slam_api_processor', true),
  ('KINK', 'https://api.kink.nl/static/now-playing.json', 'https://icecast.samuelvaneck.com/kink.mp3', 'kink_api_processor', true),
  ('100% NL', 'https://api.100p.nl/api/live?brand=100pnl', 'https://icecast.samuelvaneck.com/100pnl.mp3', 'slam_api_processor', true);
