import { useState } from "react";
import { createVaultItem } from "../api/vaultService";

export function VaultForm({ onSuccess }: { onSuccess: () => void }) {
  const [form, setForm] = useState({
    title: "",
    username: "",
    password: "",
    notes: "",
  });
  const [error, setError] = useState("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await createVaultItem(form);
      onSuccess();
      setForm({ title: "", username: "", password: "", notes: "" });
    } catch (err: any) {
      setError("Failed to create vault item.");
      console.error(err);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4 p-4 border rounded-xl shadow-sm bg-white mb-6">
      {error && <p className="text-red-500">{error}</p>}
      <input
        name="title"
        value={form.title}
        onChange={handleChange}
        placeholder="Title"
        className="w-full p-2 border rounded"
        required
      />
      <input
        name="username"
        value={form.username}
        onChange={handleChange}
        placeholder="Username"
        className="w-full p-2 border rounded"
        required
      />
      <input
        name="password"
        value={form.password}
        onChange={handleChange}
        placeholder="Password"
        type="password"
        className="w-full p-2 border rounded"
        required
      />
      <textarea
        name="notes"
        value={form.notes}
        onChange={handleChange}
        placeholder="Notes"
        className="w-full p-2 border rounded"
      />
      <button type="submit" className="px-4 py-2 bg-blue-500 text-white rounded">
        Add Item
      </button>
    </form>
  );
}
