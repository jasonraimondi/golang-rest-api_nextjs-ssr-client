import { Formik, FormikActions, FormikProps } from "formik";
import React from "react";
import { addTagsToPhoto } from "../lib/services/api/photos";
import { SubmitButton } from "./forms/button";
import { TextInput } from "./forms/text";

type TagInputs = {
  tags: string;
}

export const EditTags = ({ photoId, afterSave }: { photoId: string, afterSave(): void }) => {
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
    const tags = values.tags.split(", ");
    const errorMessage: string|null = await addTagsToPhoto(photoId, tags);
    if (errorMessage) setStatus(errorMessage);
    setSubmitting(false);
    afterSave()
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
    }: FormikProps<TagInputs>) => <form onSubmit={handleSubmit}>
      <p>{status ? status : null}</p>
      <TextInput type="text"
                 label="Add Tags"
                 name="tags"
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
  </Formik>;
};
