import Image from "next/image";
import HelloTest from "../../components/hello-test";

export default function Home() {
  return (
    <main className="min-h-screen p-8">
      <h1 className="text-3xl font-bold mb-6">BudSafe Dashboard</h1>
      <HelloTest />
      {/* Your other components */}
    </main>
  );
}

