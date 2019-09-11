import React, { useCallback, useEffect, useState } from "react";
import { useDropzone } from "react-dropzone";

import "@/components/forms/file_drop_zone.css";

type MyFile = File & {
  preview?: string;
}

export function FileDropZone({ values, setFiles }: any) {
  const handleAcceptedFiles = (acceptedFiles: File[]) => {
    setFiles(acceptedFiles.map(file => Object.assign(file, {
      preview: URL.createObjectURL(file)
    })));
  };
  const onDrop = useCallback(handleAcceptedFiles, []);
  const { getRootProps, getInputProps, isDragAccept, isDragActive, isDragReject } = useDropzone({ onDrop, accept: 'image/*' });

  useEffect(() => () => {
    // Make sure to revoke the data uris to avoid memory leaks
    values.files.forEach((file: any) => URL.revokeObjectURL(file.preview));
  }, [values.files]);

  return (
    <div id="file-upload-dropzone" {...getRootProps()}>
      <div className="inputs">
        <input {...getInputProps()} />
        <p>
          {isDragAccept && ("All files will be accepted")}
          {isDragReject && ("Some files will be rejected")}
          {!isDragActive && ("Drop some files here ...")}
        </p>
      </div>
      <div className="previews">
        {values.files.map((file: MyFile) => <Thumb key={file.name} file={file}/>)}
      </div>
    </div>
  );
}

function Thumb({ file }: { file?: any }) {
  const [loading] = useState(true);

  console.log(file.preview);

  if (!file) {
    return null;
  }

  if (loading && false) {
    return <p>loading...</p>;
  }

  return (<img src={file.preview}
               alt={file.name}
               className="img-thumbnail mt-2"
               height={200}
               width={200}/>);
}