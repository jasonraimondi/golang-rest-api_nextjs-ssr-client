import React, { useCallback, useEffect, useState } from "react";
import { useDropzone } from "react-dropzone";
import "./file_drop_zone.css";

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
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop, accept: 'image/*' });
  let inputMessage = "Drop the files here ...";
  if (isDragActive) {
    inputMessage = "Drag 'n' drop some files here, or click to select files";
  }

  useEffect(() => () => {
    // Make sure to revoke the data uris to avoid memory leaks
    values.files.forEach((file: any) => URL.revokeObjectURL(file.preview));
  }, [values.files]);

  return (
    <div id="file-upload-dropzone" {...getRootProps()}>
      <input {...getInputProps()} />
      <p>{inputMessage}</p>
      <div>
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