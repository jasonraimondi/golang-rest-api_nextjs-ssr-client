import { get } from "../../rest_client";

export async function listPhotos(userId: string, page: number, itemsPerPage: number) {
  const inputs = {
    userId,
    page,
    itemsPerPage,
  };
  const res: any = await get("/list_photos", inputs);
  if (res.error) {
    return res.error;
  }
  res.data.Data = res.data.Data.map((photo: any) => ({
    id: photo.ID,
    fileName: photo.FileName,
    relativeURL: photo.RelativeURL,
    sha256: photo.SHA256,
    mimeType: photo.MimeType,
    fileSize: photo.FileSize,
    width: photo.Width,
    height: photo.Height,
    userID: photo.UserId,
    createdAt: photo.CreatedAt,
    modifiedAt: photo.ModifiedAt,
  } as Photo));
  return res;
}

export interface Photo {
  id: string;
  fileName: string;
  relativeURL: string;
  sha256: string;
  mimeType: string;
  fileSize: number;
  width: NullInt;
  height: NullInt;
  userID: string;
  createdAt: number;
  modifiedAt: NullInt;
}

export interface NullInt {
  int64: number;
  valid: boolean;
}
