import React, { useCallback } from "react";
import { useDropzone } from "react-dropzone";

export function MyDropzone({ values, setFiles }: any) {
  const handleAcceptedFiles = (acceptedFiles: File[]) => {
    console.log(acceptedFiles)
    setFiles(acceptedFiles);
  };
  const onDrop = useCallback(handleAcceptedFiles, []);
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });
  return (
    <div {...getRootProps()}>
      <input {...getInputProps()} />
      {
        isDragActive ?
          <p>Drop the files here ...</p> :
          <p>Drag 'n' drop some files here, or click to select files</p>
      }
      {values.files.map((file: any, i: any) => <Thumb key={i} file={file} />)}
    </div>
  );
}

class Thumb extends React.Component<{key: number, file: File}> {
  state = {
    loading: false,
    thumb: undefined,
  };

  componentWillReceiveProps(nextProps: any) {
    if (!nextProps.file) { return; }

    this.setState({ loading: true }, () => {
      let reader = new FileReader();

      reader.onloadend = () => {
        this.setState({ loading: false, thumb: reader.result });
      };

      reader.readAsDataURL(nextProps.file);
    });
  }

  render() {
    const { file }: any = this.props;
    const { loading, thumb } = this.state;

    if (!file) { return null; }

    if (loading) { return <p>loading...</p>; }

    return (<img src={thumb}
                 alt={file.name}
                 className="img-thumbnail mt-2"
                 height={200}
                 width={200} />);
  }
}