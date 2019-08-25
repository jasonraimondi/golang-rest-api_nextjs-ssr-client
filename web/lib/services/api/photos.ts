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

export async function addTagsToPhoto(photoId: string, tags: string[]) {
  const data = new URLSearchParams();
  tags.forEach(tag => data.append("tags[]", tag));
  return await post(`/admin/photos/${photoId}/tags`, data);
}
export interface Photo {
  id: string;
  fileName: string;
  relativeURL: string;
  sha256: string;
  mimeType: string;
  tags: string[];
  fileSize: number;
  description: NullString;
  width: NullInt;
  height: NullInt;
  userID: string;
  createdAt: number;
  modifiedAt: NullInt;
}

export const ToPhoto = (photo: any) => {
  console.log(photo)
  return {
    id: photo.ID,
    tags: photo.Tags ? photo.Tags.map((tag: any) => tag.Name).sort() : [],
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
