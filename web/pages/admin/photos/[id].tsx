import { NextPageContext } from "next";
import Router from "next/router";
import React, { useState } from "react";
import { EditTags } from "../../../components/edit_tags";
import { defaultLayout } from "../../../components/layouts/default";
import { Tag } from "../../../components/tag";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { APP_ROUTES } from "../../../lib/routes";
import { addTagsToPhoto, getPhoto, Photo, PHOTO_BASE_PATH, removeTagFromPhoto } from "../../../lib/services/api/photos";

type Props = {
  photo: Photo;
} & AuthProps

function Page({ photo }: Props) {
  const [tags, setTags] = useState(photo.Tags);

  const fooBar = async (photo: Photo) => {
    const res = await addTagsToPhoto(photo.ID, ["one", "two", "dumber", "and hello"]);
    console.log({ res });
  };

  const handleRemoveTag = async (photoId: string, tagId: number) => {
    const res: any = await removeTagFromPhoto(photoId, tagId);
    if (res.status == 202) {
      setTags(tags.filter(tag => tag.ID !== tagId));
    }
  };

  const tagList = tags.length ? tags.map(tag => {
    return <Tag tag={tag} handleRemoveTag={() => handleRemoveTag(photo.ID, tag.ID)} key={tag.ID}/>;
  }) : "no tags";

  return <div className="container mx-auto max-w-sm">
    <img width={420} src={PHOTO_BASE_PATH + photo.RelativeURL} alt={photo.Description.string}
         title={photo.Description.string}/>
    <p>FileSize: {photo.FileSize}</p>
    <p>MimeType: {photo.MimeType}</p>
    <p>FileName: {photo.FileName}</p>
    <div>Tags: {tagList}</div>
    <EditTags photoId={photo.ID}
              afterSave={() => Router.push(APP_ROUTES.admin.photos.show.create({ photoId: photo.ID }))}
    />
    <button onClick={() => fooBar(photo)}>Create Tags</button>
  </div>;
}

Page.getInitialProps = async ({ query }: NextPageContext) => {
  const id: any = query["id"];
  const photo = await getPhoto(id);
  return { photo };
};

export default privateRoute(defaultLayout(Page));
