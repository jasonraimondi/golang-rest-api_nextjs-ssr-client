import React from "react";

export interface TextFields {
  type: "text" | "password" | "email";
  label: string;
  name: string;
  value?: string;
  handleInputChange: (e) => void;
  required?: boolean;
}

export function TextInput({ label, name, value, handleInputChange, required, type }: TextFields) {
  return <label className="block">
    <span className="block">
      {label}:
    </span>
    <input className="border-solid border-2 border-gray-600 rounded w-full"
           type={type} name={name} onChange={handleInputChange} value={value} required={required}/>
  </label>;
}