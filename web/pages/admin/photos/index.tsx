import Head from "next/head";
import React, { Component } from "react";
import { defaultLayout } from "../../../components/layouts/default";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { APP_ROUTES } from "../../../lib/routes";
import { listPhotosForUser, Photo, PHOTO_BASE_PATH } from "../../../lib/services/api/photos";

type Props = {
  photos: any
}

class Page extends Component<Props & AuthProps> {
  get photos() {
    if (!this.props.photos) return;
    return this.props.photos.map((photo: Photo) => {
      const photoSrc = `${PHOTO_BASE_PATH}${photo.RelativeURL}`;
      const photoId = photo.ID;
      return <li key={photo.ID}>
        <a href={APP_ROUTES.admin.photos.show.create({ photoId })}>
          <img className="max-w-xs" src={photoSrc}/>
        </a>
        <p>{photo.FileName}</p>
        <p>Tags: {photo.TagList}</p>
      </li>;
    });
  }

  static async getInitialProps({ auth }: AuthProps) {
    const res: any = await listPhotosForUser(auth.user.id, 1, 250);
    return { auth, photos: res };
  }

  render() {
    return <>
      <Head>
        <title>My Photos</title>
      </Head>
      <ul>
        {this.photos}
      </ul>
    </>;
  }
}

export default privateRoute(defaultLayout(Page));
