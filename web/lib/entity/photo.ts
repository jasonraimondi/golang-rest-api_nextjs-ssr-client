import { App } from "./app";
import { NullInt, NullString } from "./base";
import { Tag } from "./tag";

export interface Photo {
  ID: string;
  FileName: string;
  RelativeURL: string;
  SHA256: string;
  MimeType: string;
  App?: App;
  Tags: Tag[];
  TagList: string;
  FileSize: number;
  FileSizeHuman: string;
  Description: NullString;
  Width: NullInt;
  Height: NullInt;
  UserID: string;
  CreatedAt: number;
  ModifiedAt: NullInt;
}

export const ToPhoto = (data: any) => {
  const photo: Photo = data;
  photo.App = photo.App ? photo.App : undefined;

  photo.Tags = photo.Tags ? photo.Tags : [];
  photo.Tags = photo.Tags.sort(sortTagByName);
  photo.TagList = photo.Tags.map(tag => tag.Name).join(", ");

  photo.FileSizeHuman = formatSizeUnits(photo.FileSize);

  return photo;
};

function formatSizeUnits(bytes: number): string {
  let result: string;
  if (bytes >= 1073741824) {
    result = (bytes / 1073741824).toFixed(2) + " GB";
  } else if (bytes >= 1048576) {
    result = (bytes / 1048576).toFixed(2) + " MB";
  } else if (bytes >= 1024) {
    result = (bytes / 1024).toFixed(2) + " KB";
  } else if (bytes > 1) {
    result = bytes + " bytes";
  } else if (bytes == 1) {
    result = bytes + " byte";
  } else {
    result = "0 bytes";
  }
  return result;
}

export const sortTagByName = (a: any, b: any) => {
  a = a.Name.toUpperCase();
  b = b.Name.toUpperCase();
  if (a < b) {
    return -1;
  }
  if (a > b) {
    return 1;
  }

  // names must be equal
  return 0;
};
