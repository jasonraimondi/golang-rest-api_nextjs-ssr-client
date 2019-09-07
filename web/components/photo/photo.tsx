import Link from "next/link";
import React from "react";
import { Photo } from "../../lib/entity/photo";
import { PHOTO_BASE_PATH } from "../../lib/services/api/photos";

type Props = {
  photo: Photo,
  href: (obj: any) => string,
};

export function SinglePhoto({ photo, href }: Props) {
  const photoSrc = `${PHOTO_BASE_PATH}${photo.RelativeURL}`;
  const photoId = photo.ID;
  const appId = photo.App ? photo.App.ID : "unknown";
  const appSlug = photo.App ? photo.App.Name : "unknown";
  return (
    <li className="border border-grey-800" key={photo.ID}>
      <Link href={href({ appId, appSlug, photoId })}>
        <a><img src={photoSrc} alt={photo.Description.String}/></a>
      </Link>
      <div className="p-1">
        <p>{photo.FileName}</p>
        <p>Tags: {photo.TagList}</p>
      </div>
    </li>
  );
}