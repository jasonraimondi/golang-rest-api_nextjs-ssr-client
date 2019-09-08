import { postMultipart } from "../rest_client";
import { API_ROUTES } from "../routes";

interface UploadFileFields {
  userId: string;
  files: File[];
}

export async function uploadFiles(bearer: string, { userId, files }: UploadFileFields) {
  const formData = new FormData();
  files.forEach(file => formData.append("files[]", file));
  return await postMultipart(API_ROUTES.photos.upload_photo.create({ userId }), formData, {
    Authorization: bearer,
  });
}
