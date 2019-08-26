import { NextPageContext } from "next";
import React, { useState } from "react";
import { EditTags } from "../../../components/edit_tags";
import { defaultLayout } from "../../../components/layouts/default";
import { Tag } from "../../../components/tag";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { addTagsToPhoto, getPhoto, Photo, PHOTO_BASE_PATH, removeTagFromPhoto } from "../../../lib/services/api/photos";

type Props = {
  photo: Photo;
} & AuthProps

function Page({ photo }: Props) {
  const [tags, setTags] = useState(photo.Tags);

  const fooBar = async (photo: Photo) => {
    const res = await addTagsToPhoto(photo.ID, ["one", "two", "dumber", "and hello"]);
    console.log({res});
  };

  const handleRemoveTag = async (photoId: string, tagId: number) => {
    const res: any = await removeTagFromPhoto(photoId, tagId);
    if (res.status == 202) {
      setTags(tags.filter(tag => tag.ID !== tagId))
    }
  };

  const foo = tags.map(tag => {
    return <Tag tag={tag} handleRemoveTag={() => handleRemoveTag(photo.ID, tag.ID)} key={tag.ID} />;
  });

  return <>
    <img width={420} src={PHOTO_BASE_PATH + photo.RelativeURL} alt={photo.Description.string} title={photo.Description.string}/>
    <p>{photo.FileSize}</p>
    <p>{photo.MimeType}</p>
    <p>{photo.FileName}</p>
    <div>Tags: {foo}</div>
    <EditTags photoId={photo.ID}/>
    <button onClick={() => fooBar(photo)}>Create Tags</button>
  </>;
}

Page.getInitialProps = async ({ query }: NextPageContext) => {
  const id: any = query["id"];
  const photo = await getPhoto(id);
  return { photo };
};

export default privateRoute(defaultLayout(Page));
