import { Formik, FormikActions, FormikProps } from "formik";
import Dropzone from "react-dropzone";

import { SubmitButton } from "../../elements/forms/button";
import { FileInput } from "../../elements/forms/file";
import { defaultLayout } from "../../elements/layouts/default";
import { AuthProps, privateRoute } from "../../lib/auth/private_route";
import { uploadFiles } from "../../lib/services/api/upload_file";

export type PhotoUpload = {
  files: File[];
}

const dropzoneStyle = {
  width: "100%",
  height: "auto",
  borderWidth: 2,
  borderColor: "rgb(102, 102, 102)",
  borderStyle: "dashed",
  borderRadius: 5,
}

function Page({ auth }: AuthProps) {
  const initialValues: PhotoUpload = { files: [] };

  const validate = (values: PhotoUpload) => {
    let errors: any = {};

    if (!values.files) {
      errors.files = "Required";
    }

    return errors;
  };

  const onSubmit = async (values: PhotoUpload, { setSubmitting, setStatus }: FormikActions<PhotoUpload>) => {
    console.log({
      onSubmit: "hello",
      values,
    });
    await uploadFiles(auth.authorizationString, { userId: auth.user.id, files: values.files });
    if ("") setStatus("set status");
    setSubmitting(false);
    return;

  };

  return <Formik
    initialValues={initialValues}
    validate={validate}
    onSubmit={onSubmit}
  >
    {({
      values,
      status,
      touched,
      setFieldValue,
      handleChange,
      handleBlur,
      handleSubmit,
      isSubmitting,
    }: FormikProps<PhotoUpload>) => <form className="container mx-auto max-w-sm" onSubmit={handleSubmit}>
      <p>{status ? status : null}</p>


      <Dropzone style={dropzoneStyle} accept="image/*" onDrop={(acceptedFiles) => {
        // do nothing if no files
        if (acceptedFiles.length === 0) { return; }

        // on drop we add to the existing files
        setFieldValue("files", values.files.concat(acceptedFiles));
      }}>
        {({ isDragActive, isDragReject, acceptedFiles, rejectedFiles }: any) => {
          if (isDragActive) {
            return "This file is authorized";
          }

          if (isDragReject) {
            return "This file is not authorized";
          }

          if (values.files.length === 0) {
            return <p>Try dragging a file here!</p>
          }

          return values.files.map((file, i) => (<div key={i} file={file} />));
        }}
      </Dropzone>

      <FileInput label="Email"
                 name="file[]"
                 touched={!!touched.files}
                 handleBlur={handleBlur}
                 handleChange={handleChange}
                 submitting={isSubmitting}
      />
      <SubmitButton label="Upload" type="submit" disabled={isSubmitting}/>
    </form>}
  </Formik>;

}


export default privateRoute(defaultLayout(Page));
