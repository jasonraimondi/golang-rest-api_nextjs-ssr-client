import { get, post } from "../../rest_client";
import { API_ROUTES } from "../../routes";

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

export async function removeTagFromPhoto(photoId: string, tagId: number) {
  return await post(API_ROUTES.photos.remove_tag.create({ photoId, tagId: tagId.toString() }));
}

export async function removeAppFromPhoto(photoId: string, appId: number) {
  return await post(API_ROUTES.photos.remove_app.create({ photoId, appId: appId.toString() }));
}

export async function addTagsToPhoto(photoId: string, tags: string[]) {
  const data = new URLSearchParams();
  tags.forEach(tag => data.append("tags[]", tag));
  const res: any = await post(API_ROUTES.photos.add_tags.create({ photoId }), data);
  if (res.error) {
    return res.error;
  }
  if (!res.data || res.status !== 202) {
    return "Something went wrong!";
  }
}

export async function addAppsToPhoto(photoId: string, apps: string[]) {
  const data = new URLSearchParams();
  apps.forEach(tag => data.append("apps[]", tag));
  const res: any = await post(API_ROUTES.photos.add_apps.create({ photoId }), data);
  if (res.error) {
    return res.error;
  }
  if (!res.data || res.status !== 202) {
    return "Something went wrong!";
  }
}

export interface Photo {
  ID: string;
  FileName: string;
  RelativeURL: string;
  SHA256: string;
  MimeType: string;
  Apps: Tags[];
  AppList: string;
  Tags: Tags[];
  TagList: string;
  FileSize: number;
  FileSizeHuman: string;
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
  photo.Apps = photo.Apps ? photo.Apps : [];
  photo.Apps = photo.Apps.sort(sortTagByName);
  photo.AppList = photo.Apps.map(tag => tag.Name).join(", ");

  photo.Tags = photo.Tags ? photo.Tags : [];
  photo.Tags = photo.Tags.sort(sortTagByName);
  photo.TagList = photo.Tags.map(tag => tag.Name).join(", ");

  photo.FileSizeHuman = formatSizeUnits(photo.FileSize);
  return photo;
};

function formatSizeUnits(bytes: number): string {
  let result: string;
  if      (bytes >= 1073741824) { result = (bytes / 1073741824).toFixed(2) + " GB"; }
  else if (bytes >= 1048576)    { result = (bytes / 1048576).toFixed(2) + " MB"; }
  else if (bytes >= 1024)       { result = (bytes / 1024).toFixed(2) + " KB"; }
  else if (bytes > 1)           { result = bytes + " bytes"; }
  else if (bytes == 1)          { result = bytes + " byte"; }
  else                          { result = "0 bytes"; }
  return result;
}

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

