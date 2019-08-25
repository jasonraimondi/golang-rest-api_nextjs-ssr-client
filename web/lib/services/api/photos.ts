import { get } from "../../rest_client";

export const PHOTO_BASE_PATH = "http://localhost:9000/originals/";

export async function getPhoto(photoId: string) {
  const res: any = await get(`/photos/${photoId}`);
  return ToPhoto(res.data);
}

export async function listPhotos(userId: string, page: number, itemsPerPage: number) {
  const inputs = {
    page,
    itemsPerPage,
  };
  const res: any = await get(`/photos/user/${userId}`, inputs);
  if (res.error) {
    return res.error;
  }
  return res.data.records.map((photo: any) => ToPhoto(photo));
}

export interface Photo {
  id: string;
  fileName: string;
  relativeURL: string;
  sha256: string;
  mimeType: string;
  fileSize: number;
  description: NullString;
  width: NullInt;
  height: NullInt;
  userID: string;
  createdAt: number;
  modifiedAt: NullInt;
}

export const ToPhoto = (photo: any) => {
  return {
    id: photo.ID,
    fileName: photo.FileName,
    relativeURL: photo.RelativeURL,
    sha256: photo.SHA256,
    mimeType: photo.MimeType,
    fileSize: photo.FileSize,
    description: photo.Description,
    width: photo.Width,
    height: photo.Height,
    userID: photo.UserId,
    createdAt: photo.CreatedAt,
    modifiedAt: photo.ModifiedAt,
  } as Photo;
};

export interface NullString {
  string: string;
  valid: boolean;
}

export interface NullInt {
  int64: number;
  valid: boolean;
}
