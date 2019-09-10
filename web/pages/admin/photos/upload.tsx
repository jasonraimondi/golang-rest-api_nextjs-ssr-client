import { Formik, FormikActions, FormikProps } from "formik";
import { NextPage } from "next";
import Router from "next/router";

import { SubmitButton } from "../../../components/forms/button";
import { FileDropZone } from "../../../components/forms/file_drop_zone";
import { AuthProps } from "../../../components/auth/private_route";
import { adminLayout } from "../../../components/admin/admin_layout";
import { APP_ROUTES } from "../../../lib/routes";
import { uploadFiles } from "../../../lib/api/upload_file";

import "./upload.css";

export type PhotoUpload = {
  files: File[];
}

type Props = AuthProps;

const Page: NextPage<Props> = ({ auth }: Props) => {
  const initialValues: PhotoUpload = { files: [] };

  const validate = (values: PhotoUpload) => {
    let errors: any = {};

    if (!values.files) {
      errors.files = "Required";
    }

    return errors;
  };

  const onSubmit = async (values: PhotoUpload, { setSubmitting, setStatus }: FormikActions<PhotoUpload>) => {
    const res: any = await uploadFiles(auth.authorizationString, { userId: auth.user.id, files: values.files });
    if (res.error) setStatus(res.error);
    setSubmitting(false);
    if (!res.error) await Router.push(APP_ROUTES.admin.dashboard.create());
  };

  return <Formik
    initialValues={initialValues}
    validate={validate}
    onSubmit={onSubmit}
  >
    {({
      values,
      status,
      // touched,
      setFieldValue,
      // handleChange,
      // handleBlur,
      handleSubmit,
      isSubmitting,
    }: FormikProps<PhotoUpload>) => <form id="upload-photo-form" className="container mx-auto" onSubmit={handleSubmit}>
      <p>{status ? status : null}</p>

      <FileDropZone values={values} setFiles={(acceptedFiles: File[]) => {
        if (acceptedFiles.length === 0) return;
        setFieldValue("files", values.files.concat(acceptedFiles));
      }}/>

      <SubmitButton label="Upload" type="submit" disabled={isSubmitting}/>
    </form>}
  </Formik>;
};


export default adminLayout(Page);
