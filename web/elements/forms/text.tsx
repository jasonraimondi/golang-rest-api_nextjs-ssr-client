import React from "react";

interface TextInputFields {
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
}: TextInputFields) {
  return <div className="mt-2">
    <label htmlFor={name}>{label}{required ? <small> <sup>*</sup></small> : null}</label>
    <input className="border-solid border-2 border-gray-600 rounded w-full py-1 px-2"
           type={type}
           name={name}
           id={name}
           disabled={submitting || validating}
           onBlur={handleBlur}
           onChange={handleChange}
           value={value}
           required={!!required}
    />
    <span className="block text-sm">
          {error && touched && error}
    </span>
  </div>;
}
