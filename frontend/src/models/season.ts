import { Media } from './media';
import { Episode } from './episode';  // Assuming an Episode interface

export interface Season {
  MediaID: number;
  Media?: Media;  // Optional field
  Name: string;
  Overview: string;
  SeasonTMDBID: string;
  SeasonNumber: number;
  PosterPath: string;
  Episodes: Episode[];  // Array of Episodes
}