import { Formik, FormikActions, FormikProps } from "formik";
import React from "react";
import { addAppsToPhoto, addTagsToPhoto } from "../lib/services/api/photos";
import { SubmitButton } from "./forms/button";
import { TextInput } from "./forms/text";

type TagInputs = {
  names: string;
}

type NameType = "tag" | "app"

type Props = {
  type: NameType
  photoId: string;
  afterSave(): void;
};

export const EditTags = ({ type, photoId, afterSave }: Props) => {
  const initialValues = { names: "" };

  const validate = (values: TagInputs) => {
    let errors: Partial<TagInputs> = {};
    if (!values.names) {
      errors.names = "Required";
    } else if (values.names.split(", ").length === 0) {
      errors.names = `Add a ${type}`;
    }

    return errors;
  };

  const addTags = (type: NameType, photoId: string, tags: string[]) => {
    if (type === "tag") {
      return addTagsToPhoto(photoId, tags);
    }
    return addAppsToPhoto(photoId, tags);
  };

  const onSubmit = async (values: TagInputs, { setSubmitting, setStatus }: FormikActions<TagInputs>) => {
    const tags = values.names.split(", ");
    const errorMessage: string | null = await addTags(type, photoId, tags);
    if (errorMessage) setStatus(errorMessage);
    setSubmitting(false);
    afterSave();
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
      <div className="flex">
        <TextInput type="text"
                   label={type === "tag" ? "Add Tags" : "Add Apps"}
                   name={"names"}
                   touched={touched.names}
                   value={values.names}
                   error={errors.names}
                   handleBlur={handleBlur}
                   handleChange={handleChange}
                   submitting={isSubmitting}
                   required
        />
        <SubmitButton label="Submit" type="submit" disabled={isSubmitting}/>
      </div>
    </form>}
  </Formik>;
};
