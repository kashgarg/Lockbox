import { useEffect, useState } from "react";
import { getVaultItems } from "../api/vaultService";
import { VaultItem } from "../types";

export function VaultList() {
  const [items, setItems] = useState<VaultItem[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    getVaultItems()
      .then(setItems)
      .catch((err) => setError(err.message));
  }, []);

  if (error) return <p className="text-red-500">Error: {error}</p>;
  if (!items.length) return <p>Loading...</p>;

  return (
    <div className="p-4 space-y-4">
      {items.map((item) => (
        <div
          key={item.id}
          className="border rounded-xl p-4 shadow-md bg-white"
        >
          <h3 className="text-lg font-bold">{item.title}</h3>
          <p><strong>Username:</strong> {item.username}</p>
          <p><strong>Password:</strong> {item.password}</p>
          <p><strong>Notes:</strong> {item.notes}</p>
          <p className="text-sm text-gray-500">Created: {item.created_at}</p>
        </div>
      ))}
    </div>
  );
}
