import { Media } from './media';

export interface Episode {
  mediaID: number;
  media?: Media;  // Optional field
  name: string;
  overview: string;
  episodeTMDBID: string;
  episodeNumber: number;
  stillPath: string;
  runtime: number;
  seasonNumber: number;
}