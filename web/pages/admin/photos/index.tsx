import Head from "next/head";
import React, { Component } from "react";
import { defaultLayout } from "../../../elements/layouts/default";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { listPhotos, Photo } from "../../../lib/services/api/photos";

type Props = {
  photos: any
}

class Page extends Component<Props & AuthProps> {
  get photos() {
    if (!this.props.photos) return;
    return this.props.photos.Data.map((photo: Photo) => {
      const link = `http://localhost:9000/originals/${photo.relativeURL}`;
      return <li key={photo.id}>
        <a href={link}><img className="max-w-xs" src={link}/></a>
      </li>;
    });
  }

  static async getInitialProps({ auth }: AuthProps) {
    const res: any = await listPhotos(auth.user.id, 1, 25);
    return { auth, photos: res.data };
  }

  render() {
    return <>
      <Head>
        <title>My Photos</title>
      </Head>
      <ul className="flex flex-wrap justify-between">
        {this.photos}
      </ul>
    </>;
  }
}

export default privateRoute(defaultLayout(Page));
