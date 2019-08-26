import React from "react";
import { Tags } from "../lib/services/api/photos";

// export type TagProps = { tag: Tags, photoId?: string };
export type TagProps = { tag: Tags, handleRemoveTag(): void };

export function Tag({ tag, handleRemoveTag }: TagProps) {
  return <span className="rounded p-1 bg-blue-300 m-1">
    {tag.Name} <button onClick={handleRemoveTag}>&times;</button>
  </span>;
}
