export interface MediaType {
  id: number;
  name: string;
  description: string;
  deletedAt?: Date;  // Optional field
  createdAt?: Date;  // Optional field
}