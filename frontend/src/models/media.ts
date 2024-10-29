import { MediaType } from './media-type';

export interface Media {
  ID: number;
  TMDBID: number;
  Title: string;
  PosterImage: string;
  MediaTypeId: number;
  MediaType?: MediaType;  // Optional field
  DeletedAt?: Date;       // Optional field
  CreatedAt?: Date;       // Optional field
}