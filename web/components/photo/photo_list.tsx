import React from "react";
import { Photo } from "../../lib/entity/photo";

import { SinglePhoto } from "./photo";
import "./photo_list.css";

interface Props {
  photos: Photo[];
  error?: string;
  href: (obj: any) => string,
}

export function PhotoList({ photos, href, error }: Props) {
  if (!photos) return <>No Photos</>;

  if (error) {
    return <p>{error}</p>
  }

  if (photos.length === 0) {
    return <p>No photos.</p>
  }

  return <ul id="photo-list">
    {photos.map((photo: Photo) => <SinglePhoto photo={photo} href={href}/>)}
  </ul>;
}