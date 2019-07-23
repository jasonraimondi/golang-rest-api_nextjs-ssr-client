import React from "react";
import styled from "styled-components";

export interface TextFields {
  type: "text" | "password" | "email";
  label: string;
  name: string;
  value: string;
  error?: string;
  validating?: boolean;
  submitting?: boolean
  required?: boolean;
  touched?: boolean;

  handleBlur(e: React.FocusEvent<any>): void;

  handleChange(e: React.ChangeEvent<any>): void;
}

export function TextInput({
  type,
  label,
  name,
  value,
  error,
  touched,
  validating,
  submitting,
  required,
  handleBlur,
  handleChange,
}: TextFields) {
  return <Label className="block mt-3">
    <span className="block">
      {label}:
    </span>
    <input className="border-solid border-2 border-gray-600 rounded w-full py-1 px-2"
           type={type}
           name={name}
           disabled={submitting || validating}
           onBlur={handleBlur}
           onChange={handleChange}
           value={value}
           required={required}
    />
    <span className="block text-sm">
          {error && touched && error}
    </span>
  </Label>;
}

const Label = styled.label`
  &:first-child {
    margin-top: 0;
  }
`;