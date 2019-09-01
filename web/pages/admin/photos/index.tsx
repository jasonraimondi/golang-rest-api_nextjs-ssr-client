import Head from "next/head";
import React from "react";

import { defaultLayout } from "../../../components/layouts/default";
import { PhotoList } from "../../../components/photo/photo_list";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { listPhotosForUser, Photo } from "../../../lib/services/api/photos";
import { AuthToken } from "../../../lib/services/auth_token";

type Props = {
  photos: Photo[]
}

function Page({ photos }: Props & AuthProps) {
  return <>
    <Head>
      <title>My Photos</title>
    </Head>
    <PhotoList photos={photos}/>
  </>;
}

Page.getInitialProps = async (ctx: any) => {
  const auth = AuthToken.fromNext(ctx);
  const res: any = await listPhotosForUser(auth.user.id, 1, 250);
  return { auth, photos: res };
};

export default privateRoute(defaultLayout(Page));
