import { NextPage, NextPageContext } from "next";
import Router from "next/router";
import React from "react";

import { adminLayout } from "@/components/admin/admin_layout";
import { AuthProps } from "@/components/auth/private_route";
import { TextInput } from "@/components/forms/text";
import { EditPhoto } from "@/components/photo/photo_edit";
import { getPhoto, PHOTO_BASE_PATH } from "@/lib/api/photos";
import { Photo } from "@/lib/entity/photo";
import { APP_ROUTES } from "@/lib/routes";

type Props = AuthProps & {
  photo: Photo,
};

const Page: NextPage<Props> = ({ photo, auth }) => {

  const afterSave = async () => {
    await Router.push(APP_ROUTES.admin.photos.show.create({ photoId: photo.ID }))
  };

  return <div className="container mx-auto max-w-sm">
    <img width={420}
         src={PHOTO_BASE_PATH + photo.RelativeURL}
         alt={photo.Description.String}
         title={photo.Description.String}
    />
    <TextInput type="text" label="File Size" name="file-size" value={photo.FileSizeHuman} disabled={true}/>
    <TextInput type="text" label="Mime Type" name="mine-type" value={photo.MimeType} disabled={true}/>
    <TextInput type="text" label="File Name" name="file-size" value={photo.FileName} disabled={true}/>
    <EditPhoto auth={auth}
               photo={photo}
               afterSave={afterSave}
    />
  </div>;
};

Page.getInitialProps = async ({ query, auth, token }: NextPageContext & AuthProps) => {
  const id: any = query["id"];
  const photo = await getPhoto(id);
  return { photo, auth, token };
};

export default adminLayout(Page);
