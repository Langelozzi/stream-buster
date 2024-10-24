import { MediaType } from './media-type';

export interface Media {
  id: number;
  tmdbID: number;
  title: string;
  posterImage: string;
  mediaTypeId: number;
  mediaType?: MediaType;  // Optional field
  deletedAt?: Date;       // Optional field
  createdAt?: Date;       // Optional field
}