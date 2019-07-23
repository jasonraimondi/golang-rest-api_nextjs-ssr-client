import { postMultipart } from "../../rest_client";
import { catchAxiosError } from "../error";

interface UploadFileFields {
  userId: string;
  file: File;
}

export async function uploadFile(bearer: string, {userId, file}: UploadFileFields) {
  try {
    const formData = new FormData();
    formData.append("userId", userId);
    formData.append("file[]", file);
    await postMultipart("/app/upload", formData, {
      Authorization: bearer,
    });
  } catch (e) {
    return catchAxiosError(e);
  }
  return "hello sunshine my best friend";
}
