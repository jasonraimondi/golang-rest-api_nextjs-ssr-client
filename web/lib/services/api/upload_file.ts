import { postMultipart } from "../../rest_client";

interface UploadFileFields {
  userId: string;
  files: File[];
}

export async function uploadFiles(bearer: string, { userId, files }: UploadFileFields) {
  const formData = new FormData();
  formData.append("userId", userId);
  files.forEach(file => formData.append("file[]", file));
  return await postMultipart("/api/upload", formData, {
    Authorization: bearer,
  });
}
