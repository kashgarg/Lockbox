import { useState, useEffect } from "react";
import { getVaultItems, deleteVaultItem } from "../api/vaultService";
import { VaultItem } from "../types";
import VaultForm from "./VaultForm";

export function VaultList() {
  const [items, setItems] = useState<VaultItem[]>([]);
  const [editingItem, setEditingItem] = useState<VaultItem | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setLoading(true);
    getVaultItems()
      .then(setItems)
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false));
  }, []);

  const handleDelete = async (id: number) => {
    if (!window.confirm("Are you sure you want to delete this item?")) return;
    try {
      await deleteVaultItem(id);
      setItems((prev) => prev.filter((item) => item.id !== id));
    } catch (err: any) {
      setError(err.message);
    }
  };

  const handleEdit = (item: VaultItem) => {
    setEditingItem(item);
  };

  const handleUpdate = (updated: VaultItem) => {
    setItems((prev) =>
      prev.map((item) => (item.id === updated.id ? updated : item))
    );
    setEditingItem(null);
  };

  if (loading) return <p>Loading vault items...</p>;
  if (error) return <p className="text-red-500">Error: {error}</p>;
  if (!items.length) return <p>No vault items found.</p>;

  return (
    <div className="p-4">
      {editingItem && (
        <VaultForm editingItem={editingItem} onUpdate={handleUpdate} />
      )}
  
      <div className="vault-items-grid">
        {items.map((item) => (
          <div
            key={item.id}
            className="vault-item-card"
          >
            <h3 className="text-lg font-bold">{item.title}</h3>
            <p><strong>Username:</strong> {item.username}</p>
            <p><strong>Password:</strong> {item.password}</p>
            <p><strong>Notes:</strong> {item.notes}</p>
            <p className="text-sm text-gray-500">Created: {item.created_at}</p>
  
            <div className="mt-2 space-x-2">
              <button
                onClick={() => handleEdit(item)}
                className="bg-blue-500 text-white px-3 py-1 rounded"
              >
                Edit
              </button>
              <button
                onClick={() => handleDelete(item.id)}
                className="bg-red-500 text-white px-3 py-1 rounded"
              >
                Delete
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
