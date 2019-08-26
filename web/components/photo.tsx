import React from "react";
import { APP_ROUTES } from "../lib/routes";
import { Photo, PHOTO_BASE_PATH } from "../lib/services/api/photos";

export function SinglePhoto({ photo }: { photo: Photo }) {
  const photoSrc = `${PHOTO_BASE_PATH}${photo.RelativeURL}`;
  const photoId = photo.ID;
  return (
    <li className="border border-grey-800" key={photo.ID}>
      <a href={APP_ROUTES.admin.photos.show.create({ photoId })}>
        <img src={photoSrc} alt={photo.Description.string}/>
      </a>
      <div className="p-1">
        <p>{photo.FileName}</p>
        <p>Tags: {photo.TagList}</p>
      </div>
    </li>
  );
}