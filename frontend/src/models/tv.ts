import { Media } from './media';
import { Season } from './season';  // Assuming a Season interface

export interface TV {
  MediaID: number;
  Media?: Media;  // Optional field
  SeasonCount: number;
  EpisodeCount: number;
  Seasons: Season[];  // Array of Seasons
  FirstAirDate?: Date;  // Optional field
  LastAirDate?: Date;
}