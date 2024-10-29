import { Media } from './media';

export interface Episode {
  MediaID: number;
  Media?: Media;  // Optional field
  Name: string;
  Overview: string;
  EpisodeTMDBID: number;
  EpisodeNumber: number;
  StillPath: string;
  Runtime: number;
  SeasonNumber: number;
}