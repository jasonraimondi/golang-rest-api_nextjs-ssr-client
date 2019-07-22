import React from "react";
import styled from "styled-components";

export interface TextFields {
  type: "text" | "password" | "email";
  label: string;
  name: string;
  value?: string;
  handleInputChange: (e) => void;
  required?: boolean;
}

const Label = styled.label`
  &:first-child {
    margin-top: 0;
  }
`;

export function TextInput({ label, name, value, handleInputChange, required, type }: TextFields) {
  return <Label className="block mt-3">
    <span className="block">
      {label}:
    </span>
    <input className="border-solid border-2 border-gray-600 rounded w-full py-1 px-2"
           type={type} name={name} onChange={handleInputChange} value={value} required={required}/>
  </Label>;
}