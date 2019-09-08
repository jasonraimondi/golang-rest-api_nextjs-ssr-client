import { Formik, FormikActions, FormikProps } from "formik";
import React from "react";
import { updatePhoto } from "../../lib/api/photos";

import { SubmitButton } from "../forms/button";
import { TextInput } from "../forms/text";

type PhotoInputs = {
  tags: string;
  app: string;
  description: string;
}

type Props = PhotoInputs & {
  photoId: string;
  afterSave(): void;
};

export const EditPhoto = ({ photoId, afterSave, tags, description, app }: Props) => {
  const initialValues = {
    tags,
    description,
    app,
  };

  const validate = (values: PhotoInputs) => {
    let errors: Partial<PhotoInputs> = {};
    if (!values.tags) {
      errors.tags = "Required";
    } else if (!values.description) {
      errors.description = `Required`;
    }

    return errors;
  };

  const onSubmit = async (values: PhotoInputs, { setSubmitting, setStatus }: FormikActions<PhotoInputs>) => {
    const tags = values.tags.split(", ");
    // @TODO fix auth string...
    const errorMessage: string | null = await updatePhoto("authstring", photoId, tags, values.description, values.app);
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
  </Formik>;
};
