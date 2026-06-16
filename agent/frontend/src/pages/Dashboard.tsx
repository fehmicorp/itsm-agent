export default function Dashboard() {
  return (
    <div className="space-y-6">
      <h1 className="text-3xl font-bold text-blue-400">System Dashboard</h1>
      <div className="grid grid-cols-3 gap-4">
        <div className="bg-gray-800 p-6 rounded-xl border border-gray-700">
          <h3 className="text-gray-400">Status</h3>
          <p className="text-2xl font-semibold text-green-400">Healthy</p>
        </div>
        <div className="bg-gray-800 p-6 rounded-xl border border-gray-700">
          <h3 className="text-gray-400">Last Scan</h3>
          <p className="text-2xl font-semibold">2 hours ago</p>
        </div>
        <div className="bg-gray-800 p-6 rounded-xl border border-gray-700">
          <h3 className="text-gray-400">Active Agents</h3>
          <p className="text-2xl font-semibold text-blue-400">12</p>
        </div>
      </div>
    </div>
  );
}