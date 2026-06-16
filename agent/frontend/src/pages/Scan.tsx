export default function Scan() {
  return (
    <div className="text-center py-20 border-2 border-dashed border-gray-700 rounded-2xl">
      <h2 className="text-2xl font-semibold mb-4">Vulnerability Scan</h2>
      <button className="bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-lg transition-all shadow-lg">
        Start New Scan
      </button>
      <div className="mt-8">
        <div className="w-64 h-2 bg-gray-700 rounded-full mx-auto overflow-hidden">
          <div className="h-full bg-blue-500 w-1/3 animate-pulse"></div>
        </div>
        <p className="mt-2 text-sm text-gray-500">Scanning system files...</p>
      </div>
    </div>
  );
}