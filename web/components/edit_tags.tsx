import { Formik, FormikActions, FormikProps } from "formik";
import Router from "next/router";
import React from "react";
import { APP_ROUTES } from "../lib/routes";
import { addTagsToPhoto } from "../lib/services/api/photos";
import { SubmitButton } from "./forms/button";
import { TextInput } from "./forms/text";

type TagInputs = {
  tags: string;
}

export const EditTags = ({ photoId }: { photoId: string }) => {
  const initialValues = { tags: "" };

  const validate = (values: TagInputs) => {
    let errors: Partial<TagInputs> = {};
    if (!values.tags) {
      errors.tags = "Required";
    } else if (values.tags.split(", ").length === 0) {
      errors.tags = "Add a tag";
    }

    return errors;
  };

  const onSubmit = async (values: TagInputs, { setSubmitting, setStatus }: FormikActions<TagInputs>) => {
    console.log(values);
    const tags = values.tags.split(", ");
    const errorMessage: string|null = await addTagsToPhoto(photoId, tags);
    if (errorMessage) setStatus(errorMessage);
    setSubmitting(false);
    await Router.push(APP_ROUTES.admin.dashboard.create({}));
  };

  return <Formik
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
    }: FormikProps<TagInputs>) => <form className="container mx-auto max-w-sm" onSubmit={handleSubmit}>
      <p>{status ? status : null}</p>
      <TextInput type="text"
                 label="Tags"
                 name="tags"
                 touched={touched.tags}
                 value={values.tags}
                 error={errors.tags}
                 handleBlur={handleBlur}
                 handleChange={handleChange}
                 submitting={isSubmitting}
                 required
      />
      <SubmitButton label="Add" type="submit" disabled={isSubmitting}/>
    </form>}
  </Formik>;
};
