import Head from "next/head";
import React, { Component } from "react";
import { defaultLayout } from "../../../components/layouts/default";
import { SinglePhoto } from "../../../components/photo";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { listPhotosForUser, Photo } from "../../../lib/services/api/photos";
import "./index.css";

type Props = {
  photos: any
}

class Page extends Component<Props & AuthProps> {
  get photos() {
    if (!this.props.photos) return;
    return this.props.photos.map((photo: Photo) => <SinglePhoto photo={photo}/>);
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
      <ul id="photo-list">
        {this.photos}
      </ul>
    </>;
  }
}

export default privateRoute(defaultLayout(Page));
