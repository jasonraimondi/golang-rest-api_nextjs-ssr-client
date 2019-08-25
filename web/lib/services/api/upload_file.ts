import { postMultipart } from "../../rest_client";

interface UploadFileFields {
  userId: string;
  files: File[];
}

export async function uploadFiles(bearer: string, { userId, files }: UploadFileFields) {
  const formData = new FormData();
  files.forEach(file => formData.append("files[]", file));
  return await postMultipart(`/admin/photos/user/${userId}`, formData, {
    Authorization: bearer,
  });
}
