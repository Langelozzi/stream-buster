import { Media } from './media';
import { Genre } from './genre';  // Assuming a Genre interface

export interface Movie {
  MediaID: number;
  Media?: Media;   // Optional field
  Overview: string;
  BackdropImage: string;
  Genres: Genre[];  // Array of Genres
  ReleaseDate?: Date;  // Optional field
  Runtime: number;
}