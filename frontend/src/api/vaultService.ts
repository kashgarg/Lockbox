import { VaultItem } from "../types";

export async function getVaultItems() {
    const res = await fetch("http://localhost:8080/vaultitems");
    if (!res.ok) throw new Error("Failed to fetch vault items");
    return res.json();
  }

export async function createVaultItem(item: Omit<VaultItem, "id" | "created_at">) {
    const res = await fetch("http://localhost:8080/vaultitems", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(item),
    });
    if (!res.ok) {
      throw new Error("Failed to create item");
    }
    return res.json();
  }
  