import VaultForm from "../components/VaultForm";
import { VaultList } from "../components/VaultList";
import { useState } from "react";
import "./Home.css"; // 👈 Add this

export default function Home() {
  const [refresh, setRefresh] = useState(0);

  return (
    <div className="home-container">
      <h1 className="home-title">LockBox</h1>
      <VaultForm onSuccess={() => setRefresh(refresh + 1)} />
      <VaultList key={refresh} />
    </div>
  );
}


