import { NextPage } from "next";
import React from "react";
import { defaultLayout } from "../../components/layouts/default";
import { PhotoList } from "../../components/photo/photo_list";
import { Photo } from "../../lib/entity/photo";
import { APP_ROUTES } from "../../lib/routes";
import { ApiResponse } from "../../lib/services/api/api_response";
import { listPhotosForApp} from "../../lib/services/api/photos";
import { splitSlug } from "../../lib/services/slug_service";

type Props = {
  photos: ApiResponse<Photo[]>
};

const Page: NextPage<Props> = ({ photos }: Props) => {
  const [list, error] = photos;
  return <PhotoList photos={list} error={error} href={APP_ROUTES.photos.index.create}/>;
};

Page.getInitialProps = async ({ query }) => {
  const { app } = query;
  const slug = splitSlug(app.toString());
  const photos = await listPhotosForApp(slug.id, 1, 250);
  return { photos };
};

export default defaultLayout(Page);
