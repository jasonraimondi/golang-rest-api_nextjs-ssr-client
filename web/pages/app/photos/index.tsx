import React, { Component } from "react";
import styled from "styled-components";
import { defaultLayout } from "../../../elements/layouts/default";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { listPhotos, Photo } from "../../../lib/services/api/photos";

type Props = {
  photos: any
}

class Page extends Component<Props & AuthProps> {
  static async getInitialProps(props: any) {
    const { auth } = props;
    const res: any = await listPhotos(auth.user.id, 1, 5);
    return { auth, photos: res.data };
  }

  get photos() {
    return this.props.photos.Data.map((photo: Photo) => {
      const link = `http://localhost:9000/originals/${photo.relativeURL}`;
      return <li key={photo.id}>
        <a href={link}><Image src={link}/></a>
      </li>;
    });
  }

  render() {
    return <ul className="flex">
      {this.photos}
    </ul>;
  }
}

const Image = styled.img`
  height: 150px;
`;

export default privateRoute(defaultLayout(Page));
