import React from "react";

import { Photo } from "../../lib/services/api/photos";
import { SinglePhoto } from "../photo";
import "./photo_list.css";

export function PhotoList({ photos }: { photos: Photo[] }) {
  if (!photos) return <>No Photos</>;

  return <ul id="photo-list">
    {photos.map((photo: Photo) => <SinglePhoto photo={photo}/>)}
  </ul>;
}