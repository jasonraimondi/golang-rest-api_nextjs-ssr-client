import { Formik, FormikActions, FormikProps } from "formik";
import Router from "next/router";

import { SubmitButton } from "../../../components/forms/button";
import { MyDropzone } from "../../../components/forms/my_dropzone";
import { defaultLayout } from "../../../components/layouts/default";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
import { APP_ROUTES } from "../../../lib/routes";
import { uploadFiles } from "../../../lib/services/api/upload_file";

export type PhotoUpload = {
  files: File[];
}

type Props = AuthProps;

function Page({ auth }: Props) {
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
    if (!res.error) await Router.push(APP_ROUTES.admin.dashboard.create({}));
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
    }: FormikProps<PhotoUpload>) => <form className="container mx-auto max-w-sm" onSubmit={handleSubmit}>
      <p>{status ? status : null}</p>

      <MyDropzone values={values} setFiles={(acceptedFiles: File[]) => {
        if (acceptedFiles.length === 0) return;
        setFieldValue("files", values.files.concat(acceptedFiles));
      }}/>

      <SubmitButton label="Upload" type="submit" disabled={isSubmitting}/>
    </form>}
  </Formik>;

}


export default privateRoute(defaultLayout(Page));
