import { Media } from "./media";

export interface CurrentlyWatching {
  UserID?: number;
  User?: string | null;
  MediaId?: number;
  Media?: Media;
  EpisodeNumber?: number;
  SeasonNumber?: number;
  UpdatedAt?: string;
  DeletedAt?: string | null;
  CreatedAt?: string;
}
