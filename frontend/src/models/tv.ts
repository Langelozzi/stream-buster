import { Media } from './media';
import { Season } from './season';  // Assuming a Season interface

export interface TV {
  mediaID: number;
  media?: Media;  // Optional field
  overview: string;
  seasonCount: number;
  episodeCount: number;
  seasons: Season[];  // Array of Seasons
  firstAirDate?: Date;  // Optional field
}