import { VaultList } from "../components/VaultList";

export default function Home() {
  return (
    <div className="max-w-xl mx-auto mt-8">
      <h1 className="text-2xl font-bold mb-4">LockBox</h1>
      <VaultList />
    </div>
  );
}