import { useState } from "react";
import { SubmitButton } from "../../elements/forms/button";
import { defaultLayout } from "../../elements/layouts/default";
import { AuthService } from "../../lib/auth/auth_service";
import { privateRoute } from "../../lib/auth/private_route";
import { uploadFile } from "../../lib/services/api/upload_file";

function Page({auth}: { auth: AuthService }) {
  const [file, setFile] = useState<File | null>(null);

  const onFormSubmit = async (e) => {
    e.preventDefault();
    if (!file) {
      alert("Select a file");
      return;
    }
    await uploadFile(auth.authorizationString, {userId: auth.user.id, file});
  };

  const onChange = (e) => {
    setFile(e.target.files[0]);
  };

  return (
    <form onSubmit={onFormSubmit}>
      <h1>File Upload</h1>
      <input type="file" onChange={onChange} multiple/>
      <SubmitButton label="Upload"/>
    </form>
  );
}


export default privateRoute(defaultLayout(Page));
