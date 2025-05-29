import { useEffect, useState } from "react";
import { createVaultItem, updateVaultItem } from "../api/vaultService";
import { VaultItem } from "../types";

type Props = {
  onSuccess?: (item: VaultItem) => void;
  editingItem?: VaultItem;
  onUpdate?: (item: VaultItem) => void;
};

export default function VaultForm({ onSuccess, editingItem, onUpdate }: Props) {
  const [form, setForm] = useState({
    title: "",
    username: "",
    password: "",
    notes: "",
  });

  useEffect(() => {
    if (editingItem) {
      setForm({
        title: editingItem.title,
        username: editingItem.username,
        password: editingItem.password,
        notes: editingItem.notes,
      });
    }
  }, [editingItem]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (editingItem) {
        const updated = await updateVaultItem(editingItem.id, form);
        onUpdate?.(updated);
      } else {
        const created = await createVaultItem(form);
        onSuccess?.(created);
        setForm({ title: "", username: "", password: "", notes: "" });
      }
    } catch (err) {
      alert("Error saving item");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-3 bg-white p-4 rounded shadow">
      <h2 className="text-xl font-semibold">
        {editingItem ? "Edit Vault Item" : "Add Vault Item"}
      </h2>
      <input
        className="border p-2 w-full"
        name="title"
        placeholder="Title"
        value={form.title}
        onChange={handleChange}
        required
      />
      <input
        className="border p-2 w-full"
        name="username"
        placeholder="Username"
        value={form.username}
        onChange={handleChange}
        required
      />
      <input
      type="password"
      className="border p-2 w-full"
      name="password"
      placeholder="Password"
      value={form.password}
      onChange={handleChange}
      required
      />
      <textarea
        className="border p-2 w-full"
        name="notes"
        placeholder="Notes"
        value={form.notes}
        onChange={handleChange}
      />
      <button
        type="submit"
        className="bg-green-600 text-white px-4 py-2 rounded"
      >
        {editingItem ? "Update" : "Create"}
      </button>
    </form>
  );
}

