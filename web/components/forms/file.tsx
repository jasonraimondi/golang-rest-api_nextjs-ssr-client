import React from "react";
import "./file.css";

interface FileInputFields {
  label: string;
  name: string;
  error?: string;
  touched?: boolean;
  validating?: boolean;
  submitting?: boolean

  handleBlur(e: React.FocusEvent<any>): void;

  handleChange(e: React.ChangeEvent<any>): void;
}

export function FileInput({
  label,
  name,
  error,
  touched,
  validating,
  submitting,
  handleBlur,
  handleChange,
}: FileInputFields) {
  return <label className="file-label block mt-3">
    <span className="block">
      {label}:
    </span>
    <input className="border-solid border-2 border-gray-600 rounded w-full py-1 px-2"
           type="file"
           name={name}
           disabled={submitting || validating}
           onBlur={handleBlur}
           onChange={handleChange}
           required
           multiple
    />
    <span className="block text-sm">
          {error && touched && error}
    </span>
  </label>;
}
