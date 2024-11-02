import { Genre } from './genre';
import { MediaType } from './media-type';

export interface Media {
  ID: number;
  TMDBID: number;
  Title: string;
  Overview: string;
  PosterImage: string;
  MediaTypeId: number;
  MediaType?: MediaType;  // Optional field
  Genres?: Genre[];
  DeletedAt?: Date;       // Optional field
  CreatedAt?: Date;       // Optional field
}