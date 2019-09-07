import React from "react";
import { Tag } from "../lib/entity/tag";

export type TagProps = { tag: Tag, handleRemoveTag(): void };

export function Tag({ tag, handleRemoveTag }: TagProps) {
  return <span className="rounded p-1 m-1 bg-blue-300 ">
    {tag.Name} <button onClick={handleRemoveTag}>&times;</button>
  </span>;
}
