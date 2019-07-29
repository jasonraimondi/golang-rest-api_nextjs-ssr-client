import { Formik, FormikActions, FormikProps } from "formik";
import Router from "next/router";

import { SubmitButton } from "../../../elements/forms/button";
import { MyDropzone } from "../../../elements/forms/my_dropzone";
import { defaultLayout } from "../../../elements/layouts/default";
import { AuthProps, privateRoute } from "../../../lib/auth/private_route";
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
    if (!res.error) Router.push("/app/dashboard");
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

      <pre><code>{JSON.stringify(values)}</code></pre>

      <MyDropzone values={values} setFiles={(acceptedFiles: File[]) => {
        // do nothing if no files
        if (acceptedFiles.length === 0) {
          return;
        }
        // on drop we add to the existing files
        setFieldValue("files", values.files.concat(acceptedFiles));
      }}/>

      <SubmitButton label="Upload" type="submit" disabled={isSubmitting}/>
    </form>}
  </Formik>;

}


export default privateRoute(defaultLayout(Page));
