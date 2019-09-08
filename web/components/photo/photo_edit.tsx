import { Formik, FormikActions, FormikProps } from "formik";
import React, { useState } from "react";
import { removeTagFromPhoto, updatePhoto } from "../../lib/api/photos";
import { Photo } from "../../lib/entity/photo";
import { AuthToken } from "../../lib/services/auth_token";

import { SubmitButton } from "../forms/button";
import { TextInput } from "../forms/text";
import { Tag } from "../tag";

type PhotoInputs = {
  tags: string;
  app: string;
  description: string;
}

type Props = {
  auth: AuthToken
  photo: Photo;
  afterSave(): void;
};

export const EditPhoto = ({ photo, auth, afterSave }: Props) => {

  const initialValues = {
    description: photo.Description.String,
    app: photo.App ? photo.App.Name : "",
    tags: "",
  };

  const validate = (values: PhotoInputs) => {
    let errors: Partial<PhotoInputs> = {};
    // if (!values.tags) {
    //   errors.tags = "Required";
    // } else if (!values.description) {
    //   errors.description = `Required`;
    // }
    console.log(values);
    return errors;
  };

  const onSubmit = async (values: PhotoInputs, { setSubmitting, setStatus }: FormikActions<PhotoInputs>) => {
    const tags = values.tags.split(", ");
    const errorMessage: string | null = await updatePhoto(auth.authorizationString, photo.ID, tags, values.description, values.app);
    if (errorMessage) setStatus(errorMessage);
    setSubmitting(false);
    afterSave();
  };

  const [tags, setTags] = useState(photo.Tags);

  const handleRemoveTag = async (photoId: string, tagId: number) => {
    const res: any = await removeTagFromPhoto(auth.authorizationString, photoId, tagId);
    if (res.status == 202) {
      setTags(tags.filter(tag => tag.ID !== tagId));
    }
  };

  return <>
    <p className="pt-4"><strong>Tags:</strong><br /> {tags.length ? tags.map(tag => {
      return <Tag tag={tag} handleRemoveTag={() => handleRemoveTag(photo.ID, tag.ID)} key={tag.ID}/>;
    }) : "no tags"}</p>
    <Formik
      initialValues={initialValues}
      validate={validate}
      onSubmit={onSubmit}
    >
      {({
        values,
        status,
        errors,
        touched,
        handleChange,
        handleBlur,
        handleSubmit,
        isSubmitting,
      }: FormikProps<PhotoInputs>) => <form onSubmit={handleSubmit}>
        <p>{status ? status : null}</p>
        <TextInput type="text"
                   label={"Description"}
                   name={"description"}
                   touched={touched.description}
                   value={values.description}
                   error={errors.description}
                   handleBlur={handleBlur}
                   handleChange={handleChange}
                   submitting={isSubmitting}
                   required
        />
        <TextInput type="text"
                   label={"App"}
                   name={"app"}
                   touched={touched.app}
                   value={values.app}
                   error={errors.app}
                   handleBlur={handleBlur}
                   handleChange={handleChange}
                   submitting={isSubmitting}
                   required
        />
        <TextInput type="text"
                   label={"Add Tags"}
                   name={"tags"}
                   touched={touched.tags}
                   value={values.tags}
                   error={errors.tags}
                   handleBlur={handleBlur}
                   handleChange={handleChange}
                   submitting={isSubmitting}
                   required
        />
        <SubmitButton label="Submit" type="submit" disabled={isSubmitting}/>
      </form>}
    </Formik>
    </>;
};
