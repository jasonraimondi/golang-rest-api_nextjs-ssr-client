import { NextPageContext } from "next";
import React from "react";
import { defaultLayout } from "../../../components/layouts/default";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { getPhoto, Photo, PHOTO_BASE_PATH } from "../../../lib/services/api/photos";

type Props = {
  photo: Photo;
} & AuthProps

function Page({ photo }: Props) {
  return <>
    <img src={PHOTO_BASE_PATH + photo.relativeURL} alt={photo.description.string} title={photo.description.string}/>
    <p>{photo.fileSize}</p>
    <p>{photo.mimeType}</p>
    <p>{photo.fileName}</p>
  </>;
}

Page.getInitialProps = async ({ query }: NextPageContext) => {
  const id: any = query["id"];
  const photo = await getPhoto(id);
  return { photo };
};

export default privateRoute(defaultLayout(Page));
