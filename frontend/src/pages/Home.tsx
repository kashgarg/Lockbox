import { VaultList } from "../components/VaultList";
import { VaultForm } from "../components/VaultForm";
import { useState } from "react";

export default function Home() {
  const [refresh, setRefresh] = useState(0);

  return (
    <div className="max-w-xl mx-auto mt-8">
      <h1 className="text-2xl font-bold mb-4">LockBox</h1>
      <VaultForm onSuccess={() => setRefresh(refresh + 1)} />
      <VaultList key={refresh} />
    </div>
  );
}
