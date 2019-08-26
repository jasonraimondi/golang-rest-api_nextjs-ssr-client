import React from "react";

export const EditTags = (props: any) => {
  console.log(props);

  return <form onSubmit={(e) => {
    console.log('hi ya sucker');
    e.preventDefault();
  }}>

    <button type="submit">Submit</button>
  </form>;
};
