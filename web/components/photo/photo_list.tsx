import React from "react";

import { Photo } from "../../lib/services/api/photos";
import { SinglePhoto } from "./photo";
import "./photo_list.css";

interface Props {
  photos: Photo[];
  href: (obj: any) => string,
}

export function PhotoList({ photos, href }: Props) {
  if (!photos) return <>No Photos</>;

  return <ul id="photo-list">
    {photos.map((photo: Photo) => <SinglePhoto photo={photo} href={href}/>)}
  </ul>;
}