import { Media } from './media';
import { Episode } from './episode';  // Assuming an Episode interface

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