"use client";

import { gql, useQuery } from "@apollo/client";

const ALL_USERS_QUERY = gql`
  query {
    users {
      id
      firstName
      lastName
      email
      role
    }
  }
`;

export default function QueryTest() {
  const { loading, error, data } = useQuery(ALL_USERS_QUERY);

  return (
    <div className="p-4 border rounded-md bg-white shadow-sm">
      <h2 className="text-xl text-gray-400 font-semibold mb-2">
        Backend Connection Test
      </h2>
      {loading && <p className="text-gray-800">Loading...</p>}
      {error && <p className="text-red-500">Error: {error.message}</p>}
      {data && (
        <table className="min-w-full text-sm text-green-700 mt-2">
          <thead>
            <tr>
              <th className="px-2 py-1 text-left">ID</th>
              <th className="px-2 py-1 text-left">First Name</th>
              <th className="px-2 py-1 text-left">Last Name</th>
              <th className="px-2 py-1 text-left">Email</th>
              <th className="px-2 py-1 text-left">Role</th>
            </tr>
          </thead>
          <tbody>
            {data.users.map((user: any) => (
              <tr key={user.id} className="border-t">
                <td className="px-2 py-1">{user.id}</td>
                <td className="px-2 py-1">{user.firstName}</td>
                <td className="px-2 py-1">{user.lastName}</td>
                <td className="px-2 py-1">{user.email}</td>
                <td className="px-2 py-1">{user.role}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
