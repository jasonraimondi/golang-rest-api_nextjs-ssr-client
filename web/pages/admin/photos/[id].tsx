import { NextPageContext } from "next";
import React from "react";
import { EditTags } from "../../../components/edit_tags";
import { defaultLayout } from "../../../components/layouts/default";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { addTagsToPhoto, getPhoto, Photo, PHOTO_BASE_PATH, Tags } from "../../../lib/services/api/photos";

type Props = {
  photo: Photo;
} & AuthProps

export function Tag({ tag }: { tag: Tags }) {
  const handleRemoveTag = (id: number) => {
    alert(`remove this tag! ${id}`)
  };
  return <span className="rounded p-1 bg-blue-300 m-1">{tag.Name} <button onClick={() => handleRemoveTag(tag.ID)}>&times;</button></span>;
}

function Page({ photo }: Props) {
  const fooBar = async (photo: Photo) => {
    const res = await addTagsToPhoto(photo.ID, ["one", "two", "dumber", "and hello"]);
    console.log(res);
  };


  const tags = photo.Tags.map(tag => {
    return <Tag tag={tag} key={tag.ID} />;
  });

  return <>
    <img src={PHOTO_BASE_PATH + photo.RelativeURL} alt={photo.Description.string} title={photo.Description.string}/>
    <p>{photo.FileSize}</p>
    <p>{photo.MimeType}</p>
    <p>{photo.FileName}</p>
    <div>Tags: {tags}</div>
    <EditTags/>
    <button onClick={() => fooBar(photo)}>Create Tags</button>
  </>;
}

Page.getInitialProps = async ({ query }: NextPageContext) => {
  const id: any = query["id"];
  const photo = await getPhoto(id);
  return { photo };
};

export default privateRoute(defaultLayout(Page));
