import { postMultipart } from "../../rest_client";
import { catchAxiosError } from "../error";

interface UploadFileFields {
  userId: string;
  file: File;
}

export async function uploadFile(bearer: string, { userId, file }: UploadFileFields) {
  console.log(bearer);
  try {
    const formData = new FormData();
    formData.append("userId", userId);
    formData.append("file[]", file);
    const foo = await postMultipart("/api/upload", formData, {
      Authorization: bearer,
    });
    console.log(foo);
  } catch (e) {
    console.log("FAIL UPLOAD", e);
    return catchAxiosError(e);
  }
  return "hello sunshine my best friend";
}
