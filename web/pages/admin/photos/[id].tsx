import { NextPageContext } from "next";
import React from "react";
import { defaultLayout } from "../../../components/layouts/default";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { addTagsToPhoto, getPhoto, Photo, PHOTO_BASE_PATH } from "../../../lib/services/api/photos";

type Props = {
  photo: Photo;
} & AuthProps

export function Tag({ name }: { name: string }) {
  return <span className="rounded p-1 bg-blue-300 m-1">Hi ya {name} &times;</span>;
}

function Page({ photo }: Props) {
  const fooBar = async (photo: Photo) => {
    const res = await addTagsToPhoto(photo.id, ["one", "two", "dumber", "and hello"]);
    console.log(res);
  };


  const tags = photo.tags.map(name => {
    return <Tag name={name}/>;
  });

  return <>
    <img src={PHOTO_BASE_PATH + photo.relativeURL} alt={photo.description.string} title={photo.description.string}/>
    <p>{photo.fileSize}</p>
    <p>{photo.mimeType}</p>
    <p>{photo.fileName}</p>
    <div>Tags: {tags}</div>
    <button onClick={() => fooBar(photo)}>Create Tags</button>
  </>;
}

Page.getInitialProps = async ({ query }: NextPageContext) => {
  const id: any = query["id"];
  const photo = await getPhoto(id);
  return { photo };
};

export default privateRoute(defaultLayout(Page));
