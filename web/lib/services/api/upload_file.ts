import { postMultipart } from "../../rest_client";
import { catchAxiosError } from "../error";

interface UploadFileFields {
  userId: string;
  files: File[];
}

export async function uploadFiles(bearer: string, { userId, files }: UploadFileFields) {
  try {
    const formData = new FormData();
    formData.append("userId", userId);
    files.forEach(file => formData.append("file[]", file));
    await postMultipart("/api/upload", formData, {
      Authorization: bearer,
    });
  } catch (e) {
    return catchAxiosError(e);
  }
  return "hello sunshine my best friend";
}
