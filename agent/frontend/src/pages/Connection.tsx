export default function Connection() {
  return (
    <div className="max-w-md mx-auto bg-gray-800 p-8 rounded-xl border border-gray-700">
      <h2 className="text-xl font-bold mb-6">Server Configuration</h2>
      <div className="space-y-4">
        <div>
          <label className="block text-sm text-gray-400 mb-1">Server URL</label>
          <input className="w-full bg-gray-900 border border-gray-700 rounded p-2 focus:border-blue-500 outline-none" type="text" placeholder="https://uem.company.com" />
        </div>
        <div>
          <label className="block text-sm text-gray-400 mb-1">API Key</label>
          <input className="w-full bg-gray-900 border border-gray-700 rounded p-2 focus:border-blue-500 outline-none" type="password" placeholder="••••••••" />
        </div>
        <button className="w-full bg-green-600 py-2 rounded font-semibold mt-4 hover:bg-green-700">
          Test Connection
        </button>
      </div>
    </div>
  );
}