import { NextPageContext } from "next";
import React from "react";
import { defaultLayout } from "../../components/layouts/default";
import { PhotoList } from "../../components/photo/photo_list";
import { APP_ROUTES } from "../../lib/routes";
import { listPhotosForApp } from "../../lib/services/api/photos";

function Page({ photos }: any) {
  return <PhotoList photos={photos} href={APP_ROUTES.photos.index.create} />;
}

Page.getInitialProps = async ({ query }: NextPageContext) => {
  const { app } = query;
  const slug = splitSlug(app.toString());
  const photos = await listPhotosForApp(slug.id, 1, 250);
  return { photos };
};

export default defaultLayout(Page);

function splitSlug(app: string) {
  const slug = app.split("-");
  return {
    id: slug[0],
    slug: slug[1],
  };
}
