import { NextPageContext } from "next";
import Router from "next/router";
import React, { useState } from "react";
import { EditTags } from "../../../components/edit_tags";
import { defaultLayout } from "../../../components/layouts/default";
import { Tag } from "../../../components/tag";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { APP_ROUTES } from "../../../lib/routes";
import {
  getPhoto,
  Photo,
  PHOTO_BASE_PATH,
  removeAppFromPhoto,
  removeTagFromPhoto,
} from "../../../lib/services/api/photos";

type Props = {
  photo: Photo;
} & AuthProps

function Page({ photo }: Props) {
  const [tags, setTags] = useState(photo.Tags);
  const [apps, setApps] = useState(photo.Apps);

  const handleRemoveTag = async (photoId: string, tagId: number) => {
    const res: any = await removeTagFromPhoto(photoId, tagId);
    if (res.status == 202) {
      setTags(tags.filter(tag => tag.ID !== tagId));
    }
  };

  const handleRemoveApp = async (photoId: string, appId: number) => {
    const res: any = await removeAppFromPhoto(photoId, appId);
    if (res.status == 202) {
      setApps(apps.filter(app => app.ID !== appId));
    }
  };

  const appList = apps.length ? apps.map(app => {
    return <Tag tag={app} handleRemoveTag={() => handleRemoveApp(photo.ID, app.ID)} key={app.ID}/>;
  }) : "no tags";

  const tagList = tags.length ? tags.map(tag => {
    return <Tag tag={tag} handleRemoveTag={() => handleRemoveTag(photo.ID, tag.ID)} key={tag.ID}/>;
  }) : "no tags";

  return <div className="container mx-auto max-w-sm">
    <img width={420} src={PHOTO_BASE_PATH + photo.RelativeURL} alt={photo.Description.string}
         title={photo.Description.string}/>
    <p><strong>FileSize:</strong> {photo.FileSizeHuman}</p>
    <p><strong>MimeType:</strong> {photo.MimeType}</p>
    <p><strong>FileName:</strong> {photo.FileName}</p>
    <p><strong>Apps:</strong> {appList}</p>
    <EditTags type="app"
              photoId={photo.ID}
              afterSave={() => Router.push(APP_ROUTES.admin.photos.show.create({ photoId: photo.ID }))}
    />
    <p><strong>Tags:</strong> {tagList}</p>
    <EditTags type="tag"
              photoId={photo.ID}
              afterSave={() => Router.push(APP_ROUTES.admin.photos.show.create({ photoId: photo.ID }))}
    />
  </div>;
}

Page.getInitialProps = async ({ query }: NextPageContext) => {
  const id: any = query["id"];
  const photo = await getPhoto(id);
  return { photo };
};

export default privateRoute(defaultLayout(Page));
