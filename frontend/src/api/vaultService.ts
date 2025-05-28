export async function getVaultItems() {
    const res = await fetch("http://localhost:8080/vaultitems");
    if (!res.ok) throw new Error("Failed to fetch vault items");
    return res.json();
  }