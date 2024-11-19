import { Media } from "./media";

export interface CurrentlyWatching {
  UserID?: number;
  User?: string | null;
  MediaId?: number;
  Media?: Media;
  EpisodeNumber?: number;
  SeasonNumber?: number;
  DeletedAt?: string | null;
  CreatedAt?: string;
}
