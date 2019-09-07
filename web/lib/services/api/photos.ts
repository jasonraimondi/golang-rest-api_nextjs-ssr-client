import { ENVIRONMENT } from "../../constants";
import { Photo, ToPhoto } from "../../entity/photo";
import { get, post } from "../../rest_client";
import { API_ROUTES } from "../../routes";
import { ApiResponse } from "./api_response";

export const PHOTO_BASE_PATH = ENVIRONMENT.s3_url;

export async function getPhoto(photoId: string) {
  const res: any = await get(`/photos/${photoId}`);
  return ToPhoto(res.data);
}

export async function listPhotosForApp(appId: string, page: number, itemsPerPage: number): Promise<ApiResponse<Photo[]>> {
  const inputs = { page, itemsPerPage };
  const res: any = await get(API_ROUTES.photos.app.create({ appId }), inputs);

  if (res.data && res.data.records) {
    return [
      res.data.records.map((photo: any) => ToPhoto(photo)),
      undefined,
    ];
  }

  if (res.error) {
    return [
      [],
      res.error,
    ];
  }

  return [
    [],
    "something went wrong!",
  ];
}

export async function listPhotos(page: number, itemsPerPage: number) {
  const inputs = { page, itemsPerPage };
  const res: any = await get(`/photos`, inputs);
  if (res.error) {
    return res.error;
  }
  return res.data.records.map((photo: any) => ToPhoto(photo));
}

export async function listPhotosForTags(tags: string[], page: number, itemsPerPage: number) {
  const inputs = { page, itemsPerPage, tags };
  const res: any = await get(`/photos/tags`, inputs);
  if (res.error) {
    return res.error;
  }
  return res.data.records.map((photo: any) => ToPhoto(photo));
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

export async function updatePhoto(photoId: string, tags: string[], description: string, app: string) {
  const data = new URLSearchParams({ description, app });
  tags.forEach(tag => data.append("tags[]", tag));
  const res: any = await post(API_ROUTES.photos.update.create({ photoId }), data);
  if (res.error) {
    return res.error;
  }
  if (!res.data || res.status !== 202) {
    return "Something went wrong!";
  }
}

