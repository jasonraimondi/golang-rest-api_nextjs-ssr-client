import { NextPage } from "next";
import Router from "next/router";
import React, { useState } from "react";
import { TextInput } from "../../../components/forms/text";
import { adminLayout } from "../../../components/layouts/admin/admin_layout";
import { EditPhoto } from "../../../components/photo/photo_edit";
import { Tag } from "../../../components/tag";
import { Photo } from "../../../lib/entity/photo";
import { APP_ROUTES } from "../../../lib/routes";
import { getPhoto, PHOTO_BASE_PATH, removeTagFromPhoto } from "../../../lib/services/api/photos";

type Props = {
  photo: Photo,
};

const Page: NextPage<Props> = ({ photo }: Props) => {
  const [tags, setTags] = useState(photo.Tags);

  const handleRemoveTag = async (photoId: string, tagId: number) => {
    const res: any = await removeTagFromPhoto(photoId, tagId);
    if (res.status == 202) {
      setTags(tags.filter(tag => tag.ID !== tagId));
    }
  };

  const tagList = tags.length ? tags.map(tag => {
    return <Tag tag={tag} handleRemoveTag={() => handleRemoveTag(photo.ID, tag.ID)} key={tag.ID}/>;
  }) : "no tags";

  return <div className="container mx-auto max-w-sm">
    <img width={420} src={PHOTO_BASE_PATH + photo.RelativeURL} alt={photo.Description.String}
         title={photo.Description.String}/>
    <TextInput type="text" label="File Size" name="file-size" value={photo.FileSizeHuman} disabled={true}/>
    <TextInput type="text" label="Mime Type" name="mine-type" value={photo.MimeType} disabled={true}/>
    <TextInput type="text" label="File Name" name="file-size" value={photo.FileName} disabled={true}/>
    <p className="pt-4"><strong>Tags:</strong><br /> {tagList}</p>
    <EditPhoto photoId={photo.ID}
               app={photo.App ? photo.App.Name : ""}
               tags={""}
               description={photo.Description ? photo.Description.String : ""}
               afterSave={() => Router.push(APP_ROUTES.admin.photos.show.create({ photoId: photo.ID }))}
    />
  </div>;
};

Page.getInitialProps = async ({ query }) => {
  const id: any = query["id"];
  const photo = await getPhoto(id);
  return { photo };
};

export default adminLayout(Page);
