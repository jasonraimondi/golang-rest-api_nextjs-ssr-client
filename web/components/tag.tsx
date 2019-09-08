import React from "react";
import { Tag as TagEntity } from "../lib/entity/tag";

export type TagProps = { tag: TagEntity, handleRemoveTag(): void };

export function Tag({ tag, handleRemoveTag }: TagProps) {
  return <span className="rounded p-1 m-1 bg-blue-300 text-white">
    {tag.Name} <button onClick={handleRemoveTag}>&times;</button>
  </span>;
}
