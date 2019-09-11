import Head from "next/head";
import React from "react";

import { adminLayout } from "@/components/admin/admin_layout";
import { PhotoList } from "@/components/photo/photo_list";
import { Photo } from "@/lib/entity/photo";
import { APP_ROUTES } from "@/lib/routes";
import { listPhotosForUser} from "@/lib/api/photos";
import { AuthToken } from "@/lib/services/auth_token";

type Props = {
  photos: Photo[]
}

function Page({ photos }: Props) {
  return <>
    <Head>
      <title>My Photos</title>
    </Head>
    <PhotoList photos={photos} href={APP_ROUTES.admin.photos.show.create}/>
  </>;
}

Page.getInitialProps = async (ctx: any) => {
  const auth = AuthToken.fromNext(ctx);
  const res: any = await listPhotosForUser(auth.user.id, 1, 250);
  return { photos: res };
};

export default adminLayout(Page);
