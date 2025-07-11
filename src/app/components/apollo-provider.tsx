"use client";

import { ApolloProvider } from "@apollo/client";
import { getApolloClient } from "../../lib/apollo-client";
import { ReactNode } from "react";

export default function ApolloProviderWrapper({
  children,
}: {
  children: ReactNode;
}) {
  const client = getApolloClient();

  return <ApolloProvider client={client}>{children}</ApolloProvider>;
}
