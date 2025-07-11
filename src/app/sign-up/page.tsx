"use client";

import { useState } from "react";
import { getAuth, createUserWithEmailAndPassword } from "firebase/auth";
import { GraphQLClient, gql } from "graphql-request";

// --- Firebase Initialization ---
// IMPORTANT: This should be done in a central file (e.g., lib/firebase.ts)
// and imported here, but for a self-contained example, it's included.
import { initializeApp, getApps } from "firebase/app";
import { getAnalytics } from "firebase/analytics";

const firebaseConfig = {
  apiKey: process.env.NEXT_PUBLIC_FIREBASE_API_KEY!,
  authDomain: process.env.NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN!,
  projectId: process.env.NEXT_PUBLIC_FIREBASE_PROJECT_ID!,
  storageBucket: process.env.NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET!,
  messagingSenderId: process.env.NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID!,
  appId: process.env.NEXT_PUBLIC_FIREBASE_APP_ID!,
  measurementId: process.env.NEXT_PUBLIC_FIREBASE_MEASUREMENT_ID!,
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);

// --- End Firebase Initialization ---

// --- GraphQL Setup ---
const GQL_ENDPOINT = "http://localhost:8080/query"; // Your backend endpoint

const CREATE_USER_MUTATION = gql`
  mutation CreateUser($input: CreateUserInput!) {
    createUser(input: $input) {
      id
      email
      firstName
      lastName
      firebaseId
      role
    }
  }
`;
// --- End GraphQL Setup ---

export default function SignUpPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");

  const [status, setStatus] = useState("idle"); // idle | loading | success | error
  const [message, setMessage] = useState("");

  const handleSignUp = async (event: React.FormEvent) => {
    event.preventDefault();
    setStatus("loading");
    setMessage("Creating user...");

    const auth = getAuth();

    // --- STEP 1: Create user in Firebase ---
    try {
      const userCredential = await createUserWithEmailAndPassword(
        auth,
        email,
        password
      );
      const firebaseUser = userCredential.user;
      setMessage("Firebase user created. Getting ID token...");

      const idToken = await firebaseUser.getIdToken();
      setMessage("ID token received. Creating user profile in our database...");

      // --- STEP 2: Create user profile in your backend ---
      const gqlClient = new GraphQLClient(GQL_ENDPOINT, {
        headers: {
          authorization: `Bearer ${idToken}`,
        },
      });

      const variables = {
        input: {
          firstName: firstName,
          lastName: lastName,
          role: "STAFF",
        },
      };

      const data = await gqlClient.request(CREATE_USER_MUTATION, variables);

      // Type assertion for the expected response structure
      const userData = data as {
        createUser: {
          id: string;
          email: string;
          firstName: string;
          lastName: string;
          firebaseId: string;
          role: string;
        };
      };

      setStatus("success");
      setMessage(
        `Success! User profile created in database. User ID: ${userData.createUser.id}`
      );
      console.log("Backend response:", userData);
    } catch (error: any) {
      setStatus("error");
      // Handle specific errors from Firebase or your backend
      const errorMessage = error.response?.errors[0]?.message || error.message;
      setMessage(`Error: ${errorMessage}`);
      console.error("Sign-up failed:", error);
    }
  };

  return (
    <div
      style={{
        maxWidth: "500px",
        margin: "50px auto",
        padding: "20px",
        border: "1px solid #ccc",
        borderRadius: "8px",
      }}>
      <h2>Sign Up & Test `createUser`</h2>
      <form onSubmit={handleSignUp}>
        <div style={{ marginBottom: "10px" }}>
          <label>First Name: </label>
          <input
            type="text"
            value={firstName}
            onChange={(e) => setFirstName(e.target.value)}
            required
            style={{ width: "100%", padding: "8px" }}
          />
        </div>
        <div style={{ marginBottom: "10px" }}>
          <label>Last Name: </label>
          <input
            type="text"
            value={lastName}
            onChange={(e) => setLastName(e.target.value)}
            required
            style={{ width: "100%", padding: "8px" }}
          />
        </div>
        <div style={{ marginBottom: "10px" }}>
          <label>Email: </label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            style={{ width: "100%", padding: "8px" }}
          />
        </div>
        <div style={{ marginBottom: "10px" }}>
          <label>Password: </label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            style={{ width: "100%", padding: "8px" }}
          />
        </div>
        <button
          type="submit"
          disabled={status === "loading"}
          style={{ padding: "10px 20px", cursor: "pointer" }}>
          {status === "loading" ? "Processing..." : "Sign Up"}
        </button>
      </form>
      {message && (
        <div
          style={{
            marginTop: "20px",
            padding: "10px",
            border: `1px solid ${status === "error" ? "red" : "green"}`,
            backgroundColor: status === "error" ? "#ffebee" : "#e8f5e9",
          }}>
          <strong>Status:</strong>
          <pre style={{ whiteSpace: "pre-wrap", wordBreak: "break-all" }}>
            {message}
          </pre>
        </div>
      )}
    </div>
  );
}
