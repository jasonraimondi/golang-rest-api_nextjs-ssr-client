import React from "react";
import { defaultLayout } from "../components/layouts/default";
import { PhotoList } from "../components/photo/photo_list";
import { APP_ROUTES } from "../lib/routes";
import { listPhotosForTags, Photo } from "../lib/services/api/photos";

function Page({ photos }: { photos: Photo[] }) {
  return <div className="w-full h-full flex items-center justify-center">
    <PhotoList photos={photos} href={APP_ROUTES.app.index.create}/>
  </div>;
}

Page.getInitialProps = async () => {
  const photos = await listPhotosForTags(["jason"], 1, 250);
  return {
    photos,
  };
};

export default defaultLayout(Page);
