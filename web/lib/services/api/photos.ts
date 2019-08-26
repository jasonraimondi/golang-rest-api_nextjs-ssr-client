import { get, post } from "../../rest_client";

export const PHOTO_BASE_PATH = "http://localhost:9000/originals/";

export async function getPhoto(photoId: string) {
  const res: any = await get(`/photos/${photoId}`);
  return ToPhoto(res.data);
}


export async function listPhotosForUser(userId: string, page: number, itemsPerPage: number) {
  const inputs = { page, itemsPerPage };
  const res: any = await get(`/photos/user/${userId}`, inputs);
  if (res.error) {
    return res.error;
  }
  return res.data.records.map((photo: any) => ToPhoto(photo));
}

export async function listPhotosForTags(tags: string[], page: number, itemsPerPage: number) {
  const inputs = { page, itemsPerPage, tags };
  const res: any = await get(`/photos/tag`, inputs);
  if (res.error) {
    return res.error;
  }
  return res.data.records.map((photo: any) => ToPhoto(photo));
}

export async function removeTagFromPhoto(photoId: string, tagId: number) {
  return await post(`/admin/photos/${photoId}/tags/${tagId}`);
}

export async function addTagsToPhoto(photoId: string, tags: string[]) {
  const data = new URLSearchParams();
  tags.forEach(tag => data.append("tags[]", tag));
  return await post(`/admin/photos/${photoId}/tags`, data);
}

export interface Photo {
  ID: string;
  FileName: string;
  RelativeURL: string;
  SHA256: string;
  MimeType: string;
  Tags: Tags[];
  TagList: string;
  FileSize: number;
  Description: NullString;
  Width: NullInt;
  Height: NullInt;
  UserID: string;
  CreatedAt: number;
  ModifiedAt: NullInt;
}

export interface Tags {
  ID: number;
  Name: string;
}

export const ToPhoto = (data: any) => {
  const photo: Photo = data;
  console.log(photo);
  photo.Tags = photo.Tags ? photo.Tags : [];
  photo.Tags = photo.Tags.sort(sortTagByName);
  photo.TagList = photo.Tags.map(tag => tag.Name).join(", ");
  return photo
};

export interface NullString {
  string: string;
  valid: boolean;
}

export interface NullInt {
  int64: number;
  valid: boolean;
}

export const sortTagByName = (a: any, b: any) => {
  a = a.Name.toUpperCase();
  b = b.Name.toUpperCase();
  if (a < b) {
    return -1;
  }
  if (a > b) {
    return 1;
  }

  // names must be equal
  return 0;
};

