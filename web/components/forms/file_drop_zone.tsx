import React, { useCallback, useState } from "react";
import { useDropzone } from "react-dropzone";
import "./file_drop_zone.css";

export function FileDropZone({ values, setFiles }: any) {
  const handleAcceptedFiles = (acceptedFiles: File[]) => {
    setFiles(acceptedFiles);
  };
  const onDrop = useCallback(handleAcceptedFiles, []);
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });
  let inputMessage = "Drop the files here ...";
  if (isDragActive) {
    inputMessage = "Drag 'n' drop some files here, or click to select files";
  }
  return (
    <div id="file-upload-dropzone" {...getRootProps()}>
      <input {...getInputProps()} />
      <p>{inputMessage}</p>
      {values.files.map((file: any, i: any) => <Thumb key={i} file={file}/>)}
    </div>
  );
}

function Thumb({ file }: { file?: File }) {
  const [loading] = useState(true);
  const [thumb] = useState(undefined);

  if (!file) {
    return null;
  }

  if (loading) {
    return <p>loading...</p>;
  }

  return (<img src={thumb}
               alt={file.name}
               className="img-thumbnail mt-2"
               height={200}
               width={200}/>);
}