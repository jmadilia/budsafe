import { ApolloClient, InMemoryCache, HttpLink } from "@apollo/client";

// Function to create Apollo Client instance
export function createApolloClient() {
  return new ApolloClient({
    link: new HttpLink({
      uri:
        process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT ||
        "http://localhost:8080/query",
    }),
    cache: new InMemoryCache(),
  });
}

// Singleton pattern for client-side
let apolloClient: ApolloClient<any> | undefined;

export function getApolloClient() {
  // Create a new client if there's none or we're on the server
  if (!apolloClient) {
    apolloClient = createApolloClient();
  }
  return apolloClient;
}
