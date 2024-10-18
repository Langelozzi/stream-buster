import { Media } from './Media';
import { Genre } from './Genre';  // Assuming a Genre interface

export interface Movie {
  mediaID: number;
  media?: Media;   // Optional field
  overview: string;
  posterPath: string;
  genres: Genre[];  // Array of Genres
  releaseDate?: Date;  // Optional field
  runtime: number;
}