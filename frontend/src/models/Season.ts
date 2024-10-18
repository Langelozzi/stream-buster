import { Media } from './Media';
import { Episode } from './Episode';  // Assuming an Episode interface

export interface Season {
  mediaID: number;
  media?: Media;  // Optional field
  name: string;
  overview: string;
  seasonTMDBID: string;
  seasonNumber: number;
  posterPath: string;
  episodes: Episode[];  // Array of Episodes
}