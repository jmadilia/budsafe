"use client";

import { gql, useQuery } from "@apollo/client";
import { useState, useEffect } from "react";

const HELLO_QUERY = gql`
  query {
    hello
  }
`;

export default function HelloTest() {
  const { loading, error, data } = useQuery(HELLO_QUERY);
  const [message, setMessage] = useState<string>("Loading...");

  useEffect(() => {
    if (loading) setMessage("Loading...");
    else if (error) setMessage(`Error: ${error.message}`);
    else if (data) setMessage(`Message from backend: ${data.hello}`);
  }, [loading, error, data]);

  return (
    <div className="p-4 border rounded-md bg-white shadow-sm">
      <h2 className="text-xl font-semibold mb-2">Backend Connection Test</h2>
      <p className={`${error ? "text-red-500" : "text-green-600"}`}>
        {message}
      </p>
    </div>
  );
}
