import React from "react";
import { Photo, PHOTO_BASE_PATH } from "../../lib/services/api/photos";

type Props = {
  photo: Photo,
  href: (obj: any) => string,
};

export function SinglePhoto({ photo, href }: Props) {
  const photoSrc = `${PHOTO_BASE_PATH}${photo.RelativeURL}`;
  const photoId = photo.ID;
  console.log(photo);
  const appId = photo.App ? photo.App.ID : "unknown";
  const appSlug = photo.App ? photo.App.Name : "unknown";
  return (
    <li className="border border-grey-800" key={photo.ID}>
      <a href={href({ appId, appSlug, photoId })}>
        <img src={photoSrc} alt={photo.Description.String}/>
      </a>
      <div className="p-1">
        <p>{photo.FileName}</p>
        <p>Tags: {photo.TagList}</p>
      </div>
    </li>
  );
}