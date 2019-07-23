import React from "react";

export interface SubmitFields {
  label: string
  type?: "submit"
  disabled?: boolean
}

export function SubmitButton({label, type, disabled}: SubmitFields) {
  if (!type) {
    type = "submit";
  }
  return <>
    <button className="p-2 mt-4 bg-blue-500 text-white rounded-full"
            type={type}
            disabled={disabled}
    >
      {label}
    </button>
  </>;
}
